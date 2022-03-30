package GoLib

import "strings"

func GetDomainUrl(url string) string {
	//https://www.baidu.com/index.php?a=1
	u := strings.Split(url, "/")
	return u[0] + "://" + u[2] + "/"
}
