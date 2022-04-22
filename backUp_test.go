package GoLib

import "testing"

func TestBackUp(t *testing.T) {
	println(len(GetBackUpDict("https://www.baidu.com")))
	for _, v := range GetBackUpDict("https://www.baidu.com") {
		println(v)
	}

}
