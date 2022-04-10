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

//判断字符串数组中是否存在字符串,flag为是否大小写敏感,1就是敏感，0就是不敏感
func StrInList(str string, strList []string, flag bool) bool {
	for _, v := range strList {
		if flag {
			str = strings.ToLower(str)
			v = strings.ToLower(v)
		}
		if str == v {
			return true
		}
	}
	return false
}
