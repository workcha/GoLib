# GoLib

向K41看起，存一些开发项目时常见的函数，封装好了以便以后直接拿来用。
# xhttp
```
兼容各种发包（JSON、文件上传、XML）,自带title、状态、请求包、返回包获取，适合用于扫描器开发的组件
```
## 函数
```
    初始化
    http := Http{Proxy: "http://127.0.0.1:8080", Header: header, TimeOut: 3}
    普通GET请求
    http.GET("http://127.0.0.1/8899.php")
    普通POST请求
    resp2 := http.POST("http://127.0.0.1/8899.php", "A=1&N2=12")
    JSON请求
    http.JsonRequest("http://127.0.0.1/8899.php", "{\"a\":\"1\"}")
    XML请求
    http.XMLRequest("http://127.0.0.1/8899.php", "<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<note>\n<to>George</to>\n<from>John</from>\n<heading>Reminder</heading>\n<body>Don't forget the meeting!</body>\n</note>")
    文件上传
    resp := http.FileUpload("file", "1.png", "http://127.0.0.1/8899.php", "image/png", []byte("I'am content"), map[string]string{"param": "1"})
    获取请求包
    println(resp.RequestPackage)
    获取返回包
	println(resp.ResponsePackage)
    ...
    //返回结果
    具体可查看测试文件或者xhttp.go
    type HttpResponse struct {
        BaseResponse	*http.Response
        Url 			string
        Status 			string
        ResponseHeader  map[string][]string
        ResponseBody 	string
        Title			string
        RequestPackage  string
        ResponsePackage	string
    }
```
### demo
```azure
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
	http := Http{Proxy: "http://127.0.0.1:8080", Header: header, TimeOut: 3}

	/*
		GET /8899.php HTTP/1.1
		Host: 127.0.0.1
		User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36
		Accept-Language: zh-CN,zh;q=0.9
		Content-Type: application/x- www-form-urlencoded
		Referer: https://item-paimai.taobao.com/pmp_item/609160317276.htm?s=pmp_detail&spm=a213x.7340941.2001.61.1aec2cb6RKlKoy
		Sec-Fetch-Mode: cors
		Sec-Fetch-Site: same-origin
		Connection: close

	*/
	http.GET("http://127.0.0.1/8899.php")
	/*
		Content-Length: 100
		Date: Sat, 02 Apr 2022 02:15:30 GMT
		Server: Apache/2.4.39 (Win32) OpenSSL/1.0.2r mod_fcgid/2.3.9a mod_log_rotate/1.02
		X-Powered-By: PHP/7.3.4
		Content-Type: text/html; charset=UTF-8
	*/
	/*
		POST /8899.php HTTP/1.1
		Host: 127.0.0.1
		User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36
		Accept-Language: zh-CN,zh;q=0.9
		Content-Type: application/x- www-form-urlencoded
		Referer: https://item-paimai.taobao.com/pmp_item/609160317276.htm?s=pmp_detail&spm=a213x.7340941.2001.61.1aec2cb6RKlKoy
		Sec-Fetch-Mode: cors
		Sec-Fetch-Site: same-origin
		Content-Length: 9
		Connection: close

		A=1&N2=12
	*/
	resp2 := http.POST("http://127.0.0.1/8899.php", "A=1&N2=12")
	println(resp2.ResponsePackage)
	/*
		POST /8899.php HTTP/1.1
		Host: 127.0.0.1
		User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36
		Accept-Language: zh-CN,zh;q=0.9
		Content-Type: application/json
		Referer: https://item-paimai.taobao.com/pmp_item/609160317276.htm?s=pmp_detail&spm=a213x.7340941.2001.61.1aec2cb6RKlKoy
		Sec-Fetch-Mode: cors
		Sec-Fetch-Site: same-origin
		Content-Length: 9
		Connection: close

		{"a":"1"}
	*/
	http.JsonRequest("http://127.0.0.1/8899.php", "{\"a\":\"1\"}")

	/*
		POST /8899.php HTTP/1.1
		Host: 127.0.0.1
		User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36
		Accept-Language: zh-CN,zh;q=0.9
		Content-Type: text/xml
		Referer: https://item-paimai.taobao.com/pmp_item/609160317276.htm?s=pmp_detail&spm=a213x.7340941.2001.61.1aec2cb6RKlKoy
		Sec-Fetch-Mode: cors
		Sec-Fetch-Site: same-origin
		Content-Length: 159
		Connection: close

		<?xml version="1.0" encoding="ISO-8859-1"?>
		<note>
		<to>George</to>
		<from>John</from>
		<heading>Reminder</heading>
		<body>Don't forget the meeting!</body>
		</note>
	*/
	http.XMLRequest("http://127.0.0.1/8899.php", "<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<note>\n<to>George</to>\n<from>John</from>\n<heading>Reminder</heading>\n<body>Don't forget the meeting!</body>\n</note>")

	/*
		POST /8899.php HTTP/1.1
		Host: 127.0.0.1
		User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36
		Content-Length: 349
		Accept-Language: zh-CN,zh;q=0.9
		Content-Type: multipart/form-data; boundary=67c9c129a6b68b717510fdff951a65ac97a9134f383bf679eb91512d77de
		Referer: https://item-paimai.taobao.com/pmp_item/609160317276.htm?s=pmp_detail&spm=a213x.7340941.2001.61.1aec2cb6RKlKoy
		Sec-Fetch-Mode: cors
		Sec-Fetch-Site: same-origin
		Connection: close

		--67c9c129a6b68b717510fdff951a65ac97a9134f383bf679eb91512d77de
		Content-Disposition: form-data; name="param"

		1
		--67c9c129a6b68b717510fdff951a65ac97a9134f383bf679eb91512d77de
		Content-Disposition: form-data; name="file"; filename="1.png"
		Content-Type: image/png

		I'am content
		--67c9c129a6b68b717510fdff951a65ac97a9134f383bf679eb91512d77de--

	*/
	//文件上传
	http.FileUpload("file", "1.png", "http://127.0.0.1/8899.php", "image/png", []byte("I'am content"), map[string]string{"param": "1"})
	
```