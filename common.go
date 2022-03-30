package GoLib

import "strings"

func GetDomainUrl(url string) string {
	//https://www.baidu.com/index.php?a=1
	u := strings.Split(url, "/")
	return u[0] + "://" + u[2] + "/"
}

// 字符串数组去重
func RemoveDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
