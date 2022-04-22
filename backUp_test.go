package GoLib

import "testing"

func TestBackUp(t *testing.T) {
	for _, v := range GetBackUpPrefix("https://www.baidu.com") {
		println(v)
	}
}
