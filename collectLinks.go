package GoLib

import (
	"github.com/anaskhan96/soup"
	"golang.org/x/net/html"
	"io"
	"strconv"
	"strings"
)

// All takes a reader object (like the one returned from http.Get())
// It returns a slice of strings representing the "href" attributes from
// anchor links found in the provided html.
// It does not close the reader passed to it.
type CollectLinksResult struct {
	Links []string
	Js    []string
	Css   []string
}

//去除重复字符串
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

//对爬取到的链接进行简单处理
//打算后期加点模板对链接去重
func dealLinks(links []string) []string {
	var result []string
	for _, link := range links {
		if link == "/" || link == "//" || strings.HasPrefix(link, "javascript") || strings.TrimSpace(link) == "" {
			continue
		}
		if strings.HasPrefix(link, "//") {
			link = "http:" + link
		}
		result = append(result, link)
	}
	return RemoveRepeatedElement(result)
}

func GetAllLinks(content string) CollectLinksResult {
	var links []string
	var js []string
	var css []string
	doc := soup.HTMLParse(content)
	//获取a href
	aLinks := doc.FindAll("a")
	for _, b := range aLinks {
		links = append(links, b.Attrs()["href"])
	}
	//script src
	jsLinks := doc.FindAll("script")
	for _, b := range jsLinks {
		js = append(js, b.Attrs()["src"])
	}
	//link href
	Links := doc.FindAll("link")
	for _, b := range Links {
		css = append(css, b.Attrs()["href"])
	}
	return CollectLinksResult{Links: dealLinks(links), Js: dealLinks(js), Css: dealLinks(css)}

}

func All(httpBody io.Reader) CollectLinksResult {
	var links []string
	var col []string
	var jsCol []string
	var js []string
	var css []string
	var cssCol []string
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return CollectLinksResult{Links: dealLinks(links), Js: dealLinks(js), Css: dealLinks(css)}
		}
		token := page.Token()
		//a标签的href属性添加，最后需要去除一些无用链接以及添加一个路径到链接里面去
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(&links, col)
				}
			}

		} else if tokenType == html.StartTagToken && token.DataAtom.String() == "form" {
			//form标签的action属性添加
			for _, attr := range token.Attr {
				if attr.Key == "action" {
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		} else if tokenType == html.StartTagToken && token.DataAtom.String() == "script" {
			//script的src标签添加
			for _, attr := range token.Attr {
				if attr.Key == "src" {
					tl := trimHash(attr.Val)
					jsCol = append(jsCol, tl)
					resolv(&js, jsCol)
				}
			}
		} else if tokenType == html.StartTagToken && token.DataAtom.String() == "link" {
			//link href
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					cssCol = append(cssCol, tl)
					resolv(&css, cssCol)
				}
			}
		}

	}
}

// trimHash slices a hash # from the link
func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

// check looks to see if a url exits in the slice.
func check(sl []string, s string) bool {
	var check bool
	for _, str := range sl {
		if str == s {
			check = true
			break
		}
	}
	return check
}

// resolv adds links to the link slice and insures that there is no repetition
// in our collection.
func resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if check(*sl, str) == false {
			*sl = append(*sl, str)
		}
	}
}
