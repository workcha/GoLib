package GoLib

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/publicsuffix"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"html"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/textproto"
	"net/url"
	"regexp"
	"strings"
	"time"
)

//初始化结构体
type Http struct {
	Proxy   string
	Header  map[string]string
	TimeOut int
	Client  *http.Client
}

//返回结果
type HttpResponse struct {
	BaseResponse    *http.Response
	Url             string
	Status          string
	ResponseHeader  map[string][]string
	ResponseBody    string
	Title           string
	RequestPackage  string
	ResponsePackage string
}

//JSON请求
func (h *Http) JsonRequest(url, body string) *HttpResponse {
	return h.httpRequest("POST", url, body, "json")
}

//xml请求
func (h *Http) XMLRequest(url, body string) *HttpResponse {
	return h.httpRequest("POST", url, body, "xml")
}

//GET请求
func (h *Http) GET(url string) *HttpResponse {
	return h.httpRequest("GET", url, "", "")
}

//POST请求
func (h *Http) POST(url string, body string) *HttpResponse {
	return h.httpRequest("POST", url, body, "")
}

//PUT请求
func (h *Http) PUT(url string, body string) *HttpResponse {
	return h.httpRequest("POST", url, body, "")
}

//HEAD请求
func (h *Http) HEAD(url string) *HttpResponse {
	return h.httpRequest("HEAD", url, "", "")
}

//正常请求
func (h *Http) httpRequest(method, url, body, tp string) *HttpResponse {
	if h.Client == nil {
		h.init()
	}
	I := bytes.NewReader([]byte(body))
	O := transform.NewReader(I, traditionalchinese.Big5.NewEncoder())
	requests, _ := http.NewRequest(method, url, O)
	for k, v := range h.Header {
		requests.Header.Add(k, v)
	}
	//json格式传输
	if tp == "json" {
		requests.Header.Add("Content-Type", "application/json")
	} else if tp == "xml" {
		//xml格式传输
		requests.Header.Add("Content-Type", "text/xml")
	} else {
		//普通格式传输
		requests.Header.Add("Content-Type", "application/x- www-form-urlencoded")
	}
	response, _ := h.Client.Do(requests)
	responseBody := getBody(response)
	return &HttpResponse{BaseResponse: response, Url: response.Request.URL.RequestURI(), Status: response.Status, ResponseHeader: response.Header, ResponseBody: responseBody, Title: getTitle(response), RequestPackage: getRequestPackage(response, body), ResponsePackage: getResponsePackage(response) + responseBody}
}

//文件上传
func (h *Http) FileUpload(fieldName, fileName, url, contentType string, fileContent []byte, params map[string]string) *HttpResponse {
	/*
		根据文件上传的底层代码重写
		req, err := NewRequest("POST", url, body)
			if err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", contentType)
			return c.Do(req)
	*/
	if h.Client == nil {
		h.init()
	}

	body := new(bytes.Buffer)
	mulWriter := multipart.NewWriter(body)
	//添加参数
	for k, v := range params {
		err := mulWriter.WriteField(k, v)
		if err != nil {
			return nil
		}
	}
	file, _ := createFormFile(fieldName, fileName, contentType, mulWriter)
	//file, err := mulWriter.CreateFormFile(fieldName, fileName)
	_, err := file.Write(fileContent)
	if err != nil {
		return nil
	}
	contentType2 := mulWriter.FormDataContentType()
	err = mulWriter.Close()
	if err != nil {
		return nil
	}

	requestBody := body.Bytes()
	requests, _ := http.NewRequest("POST", url, body)
	for k, v := range h.Header {
		requests.Header.Add(k, v)
	}
	requests.Header.Add("Content-Type", contentType2)
	response, _ := h.Client.Do(requests)
	responseBody := getBody(response)
	return &HttpResponse{BaseResponse: response, Url: response.Request.URL.RequestURI(), Status: response.Status, ResponseHeader: response.Header, ResponseBody: responseBody, Title: getTitle(response), RequestPackage: getRequestPackage(response, string(requestBody)), ResponsePackage: getResponsePackage(response) + responseBody}
}

//初始化client
func (h *Http) init() {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if h.Proxy != "" {
		proxy, _ := url.Parse(h.Proxy)
		h.Client = &http.Client{Jar: cookieJar, Timeout: time.Duration(h.TimeOut) * time.Second, Transport: &http.Transport{Proxy: http.ProxyURL(proxy), TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	} else {
		h.Client = &http.Client{Jar: cookieJar, Timeout: time.Duration(h.TimeOut) * time.Second, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	}
}

//CreateFormFile附属函数
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

//重写multipart.Writer.CreateFormFile()函数实现可以控制Content-Type内容
func createFormFile(fieldname, filename, contentType string, w *multipart.Writer) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}

//获取html-content
func getBody(response *http.Response) string {
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	return string(all)
}

func trimTitleTags(title string) string {
	// trim <title>*</title>
	titleBegin := strings.Index(title, ">")
	titleEnd := strings.Index(title, "</")
	return title[titleBegin+1 : titleEnd]
}

//获取title
func getTitle(response *http.Response) (title string) {
	body := getBody(response)
	var re = regexp.MustCompile(`(?im)<\s*title.*>(.*?)<\s*/\s*title>`)
	for _, match := range re.FindAllString(body, -1) {
		title = html.UnescapeString(trimTitleTags(match))
		break
	}
	if contentTypes, ok := response.Header["Content-Type"]; ok {
		contentType := strings.ToLower(strings.Join(contentTypes, ";"))
		if strings.Contains(contentType, "charset=gb2312") || strings.Contains(contentType, "charset=gbk") {
			titleUtf8, err := Decodegbk([]byte(title))
			if err != nil {
				return
			}
			return string(titleUtf8)
		}
	}
	return ""
}

//strng形式获取headers
func getHeaders(response *http.Response) (header string) {

	for k, v := range response.Header {
		header = header + k + ": " + strings.Join(v, ";") + "\r\n"
	}
	return header
}

//string形式获取请求包
func getRequestPackage(response *http.Response, body string) (result string) {
	request := response.Request
	result += request.Method + " " + request.URL.Path + " " + request.Proto + "\r\n"
	result += "Host: " + response.Request.Host + "\r\n"
	for k, v := range request.Header {
		result += k + ": " + strings.Join(v, ";") + "\r\n"
	}

	if body != "" {
		result += "\r\n" + body
	}
	return result
}

//获取响应包
func getResponsePackage(response *http.Response) (result string) {
	result = response.Proto + " " + response.Status + "\r\n"
	result += getHeaders(response) + "\r\n"
	return
}
