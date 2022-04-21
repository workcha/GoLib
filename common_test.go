package GoLib

import "testing"

func TestGetDomainUrl(t *testing.T) {
	url := "https://www.baidu.com/index.php?a=1"
	println(GetDomainUrl(url))

	str := RandString(20)
	println(str)
}
