package GoLib

import (
	"bytes"
	"fmt"
	"golang.org/x/net/publicsuffix"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/textproto"
	"net/url"
	"strings"
	"time"
)

type Http struct {
	Proxy   string
	Header  map[string]string
	TimeOut int
	Client  *http.Client
}

//正常请求
func (h *Http) httpRequest(method, url, body, tp string) *http.Response {
	if h.Client == nil {
		h.Init()
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
	return response
}

//JSON请求
func (h *Http) JsonRequest(url, body string) *http.Response {
	return h.httpRequest("POST", url, body, "json")
}

//xml请求
func (h *Http) XMLRequest(url, body string) *http.Response {
	return h.httpRequest("POST", url, body, "xml")
}

//GET请求
func (h *Http) GET(url string) {
	h.httpRequest("GET", url, "", "")
}

//POST请求
func (h *Http) POST(url string, body string) {
	h.httpRequest("POST", url, body, "")
}

//初始化client
func (h *Http) Init() {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if h.Proxy != "" {
		proxy, _ := url.Parse(h.Proxy)
		h.Client = &http.Client{Jar: cookieJar, Timeout: time.Duration(h.TimeOut) * time.Second, Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	} else {
		h.Client = &http.Client{Jar: cookieJar, Timeout: time.Duration(h.TimeOut) * time.Second}
	}
}

//CreateFormFile附属函数
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

//重写multipart.Writer.CreateFormFile()函数实现可以控制Content-Type内容
func CreateFormFile(fieldname, filename, contentType string, w *multipart.Writer) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}

//文件上传
func (h *Http) FileUpload(fieldName, fileName, url, contentType string, fileContent []byte, params map[string]string) *http.Response {
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
		h.Init()
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
	file, _ := CreateFormFile(fieldName, fileName, contentType, mulWriter)
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
	requests, _ := http.NewRequest("POST", url, body)
	for k, v := range h.Header {
		requests.Header.Add(k, v)
	}
	requests.Header.Add("Content-Type", contentType2)
	response, _ := h.Client.Do(requests)
	return response
}
