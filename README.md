# GoLib

向K41看起，存一些开发项目时常见的函数，封装好了以便以后直接拿来用。
#xhttp
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