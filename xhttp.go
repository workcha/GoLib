package GoLib

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/xiecat/xhttp"
	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

var client *xhttp.Client = nil

func getClient(headers map[string]string) *xhttp.Client {
	//如果已经初始化则不再进行初始化
	if client != nil && headers == nil {
		return client
	}
	options := xhttp.DefaultClientOptions()
	//设置headers
	HEADER := map[string]string{
		//"accept": "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01",
		//"accept-encoding": "gzip, deflate, br",
		"accept-language": "zh-CN,zh;q=0.9",
		"Referer":         "https://item-paimai.taobao.com/pmp_item/609160317276.htm?s=pmp_detail&spm=a213x.7340941.2001.61.1aec2cb6RKlKoy",
		"sec-fetch-mode":  "cors",
		"sec-fetch-site":  "same-origin",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36",
		//"x-requested-with": "XMLHttpRequest",
	}
	if headers != nil {
		for k := range headers {
			HEADER[k] = headers[k]
		}
	}

	options.Headers = HEADER
	options.DialTimeout = 6
	options.MaxRespBodySize = 2 << 60
	options.Proxy = "http://127.0.0.1:8080"
	// 如果要继承cookie，传入cookie jar；否则填nil。
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	// 创建client
	client, _ = xhttp.NewClient(options, cookieJar)
	return client
}

func REQUEST(url, body, method string, headers map[string]string) *xhttp.Response {
	if method == "GET" {
		return GET(url, headers)
	} else if method == "POST" {
		return POST(url, body, headers)
	}
	return nil
}

//发生GET数据包
func GET(url string, headers map[string]string) *xhttp.Response {
	c := getClient(headers)
	// 构造请求
	hr, _ := http.NewRequest("GET", url, nil)
	req := &xhttp.Request{RawRequest: hr}
	// 发起请求
	ctx := context.Background()
	resp, _ := c.Do(ctx, req)
	return resp
}

//发生POST数据包
func POST(url, body string, headers map[string]string) *xhttp.Response {
	c := getClient(headers)
	// 构造请求
	hr, _ := http.NewRequest("POST", url, strings.NewReader(body))
	req := &xhttp.Request{RawRequest: hr}
	// 发起请求
	ctx := context.Background()
	resp, _ := c.Do(ctx, req)
	return resp
}

//获取文件md5
func GetMd5(response *xhttp.Response) string {
	if response.GetStatus() == 200 {
		h := md5.New()
		h.Write(response.Body)
		return hex.EncodeToString(h.Sum(nil))
	}
	return " "
}

func trimTitleTags(title string) string {
	// trim <title>*</title>
	titleBegin := strings.Index(title, ">")
	titleEnd := strings.Index(title, "</")
	return title[titleBegin+1 : titleEnd]
}

//获取title
func GetTitle(r *xhttp.Response) (title string) {
	body := string(r.GetBody())
	var re = regexp.MustCompile(`(?im)<\s*title.*>(.*?)<\s*/\s*title>`)
	for _, match := range re.FindAllString(body, -1) {
		title = html.UnescapeString(trimTitleTags(match))
		break
	}
	// Non UTF-8
	if contentTypes, ok := r.GetHeaders()["Content-Type"]; ok {
		contentType := strings.ToLower(strings.Join(contentTypes, ";"))
		// special cases
		if strings.Contains(contentType, "charset=gb2312") || strings.Contains(contentType, "charset=gbk") {
			titleUtf8, err := Decodegbk([]byte(title))
			if err != nil {
				return
			}

			return string(titleUtf8)
		}
	}
	return
}
