package GoLib

import "testing"

func TestName(t *testing.T) {
	mysql := Mysql{UserName: "root", Password: "root", IpAddress: "127.0.0.1", Port: 3306, DbName: "Links", Charset: "utf8"}
	mysql.ConnectMysql()
	type Link struct {
		Id          int    `db:"id"`
		Url         string `db:"url"`
		Status_code int    `db:"status_code"`
		Title       string `db:"title"`
		IsUsed      int    `db:"isUsed"`
	}
	var linkResult []Link
	ok := mysql.Select(&linkResult, "SELECT * from link LIMIT 0,2")
	if ok == true {
		for _, value := range linkResult {
			println(value.Url + "\t" + value.Title)
		}
	}
}
