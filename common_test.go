package GoLib

import "testing"

func TestGetDomainUrl(t *testing.T) {
	url := "https://www.baidu.com/index.php?a=1"
	println(GetDomainUrl(url))

	str := RandString(20)
	println(str)

	println(Str2Int("123"))
	header := map[string]string{
		//"accept": "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01",
		//"accept-encoding": "gzip, deflate, br",
		"accept-language": "zh-CN,zh;q=0.9",
		"Referer":         "https://item-paimai.taobao.com/pmp_item/609160317276.htm?s=pmp_detail&spm=a213x.7340941.2001.61.1aec2cb6RKlKoy",
		"sec-fetch-mode":  "cors",
		"sec-fetch-site":  "same-origin",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36",
		//"x-requested-with": "XMLHttpRequest",
	}
	//http := Http{Proxy: "http://127.0.0.1:8080", Header: header, TimeOut: 3}
	http := Http{Header: header, TimeOut: 3}
	lines := ReadLine("1.txt")
	for _, line := range lines {
		resp := http.GET(line)
		if resp != nil {
			if LoginPageCheck(string(resp.ResponseBody)) {
				println("[*]login page: " + line + ",title: " + resp.Title)
			}
		}
	}
}
