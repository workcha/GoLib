package GoLib

import (
	"strings"
	"time"
)

//主要获取备份字典

//获取字典前缀
func getBackUpSuffix() []string {
	return []string{"gz", "sql", "bak", "log", "old", "tar.gz", "tar.tgz", "rar", "zip", "tar.bz2", "7z", "sql~", "sql.gz", "tar.7z", "tar.xz", "sql.zip"}
}

//年份的备份，一般都是5
func getYearPrefix(num int) (result []string) {
	year := time.Now().Year()
	for i := 0; i <= num; i++ {
		result = append(result, Int2Str(year-i))
	}
	year = 0
	return
}

//域名的备份
func getDomainPrefix(url string) (result []string) {
	domain := GetDomain(url)
	//ip就只添加ip即可
	if IsIp(domain) == false {
		//直接把域名丢进去的形式
		result = append(result, domain)
		domains := strings.Split(domain, ".")

		//最简单的分割加入
		for _, d := range domains {
			result = append(result, d)
		}
		//直接字符串拼接的形式
		for k, _ := range domains {
			result = append(result, strings.Join(domains[k:], "."))
			result = append(result, strings.Join(domains[k:], "-"))
			result = append(result, strings.Join(domains[k:], "_"))
		}

	} else {
		result = append(result, domain)
	}
	domain = ""
	return
}

//获取字典后缀
func getBackUpPrefix(url string) []string {
	result := []string{"bbs", "web", "www", "forum", "backup", "wwwroot", "upload", "0", "1", "123", "temp", "website", "db", "data", "code", "oa", "sysadmin", "test", "新建文件夹", "database", "user", "shopping", "sql"}
	return result
}

//获取备份字典
func GetBackUpDict(url string) (result []string) {
	suffix := getBackUpSuffix()
	prefix := getBackUpPrefix(url)
	for _, p := range prefix {
		for _, s := range suffix {
			result = append(result, p+"."+s)
		}
	}
	suffix = nil
	prefix = nil
	return
}
