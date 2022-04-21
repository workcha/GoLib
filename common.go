package GoLib

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//获取主域名url
func GetDomainUrl(url string) string {
	//https://www.baidu.com/index.php?a=1
	u := strings.Split(url, "/")
	return u[0] + "://" + u[2] + "/"
}

//从Url中获取主域名
func GetDomain(url string) string {
	//https://www.baidu.com/index.php?a=1
	return strings.Split(url, "/")[2]
}

//处理url
func DealUrl(url string) string {
	var result []string
	u := strings.Split(url, "/")
	result = append(result, u[0:3]...)
	result[0] = u[0]
	result[1] = u[1]
	result[2] = u[2]
	for i := 3; i < len(u); i++ {
		if strings.TrimSpace(u[i]) != "" {
			result = append(result, u[i])
		}
	}
	return strings.Join(result, "/")
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

//获取随机字符串
func RandString(length int) string {
	var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var strByteLen = len(strByte)
	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = strByte[r.Intn(strByteLen)]
	}

	return fmt.Sprintf("%s", bytes)
}

//判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

//字符串转int
func Str2Int(text string) int {
	d, _ := strconv.ParseInt(text, 10, 64)
	return int(d)
}
