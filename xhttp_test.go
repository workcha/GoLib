package GoLib

import (
	"testing"
)

func TestGET(t *testing.T) {
	resp := GET("https://www.baidu.com", nil)
	println(string(resp.GetBody()))
}

func TestPOST(t *testing.T) {
	resp := POST("https://www.baidu.com", "123456", nil)
	println(string(resp.GetBody()))
}

func TestREQUEST(t *testing.T) {

}

func TestGetMd5(t *testing.T) {
	resp := GET("https://www.baidu.com/favicon.ico", nil)
	print(GetMd5(resp))
}

func TestGetTitle(t *testing.T) {
	resp := GET("https://www.baidu.com/", nil)
	print(GetTitle(resp))
}
