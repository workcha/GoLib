package GoLib

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type Mysql struct {
	UserName  string
	Password  string
	IpAddress string
	Port      int
	DbName    string
	Charset   string
	db        *sqlx.DB
}

func (mysql *Mysql) ConnectMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", mysql.UserName, mysql.Password, mysql.IpAddress, mysql.Port, mysql.DbName, mysql.Charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
		os.Exit(0)
	}
	mysql.db = Db
}

//只执行，返回执行是否成功，一般用于delete、update、insert
//怕注入使用fmt进行格式化后传入
func (mysql *Mysql) Execute(sql string) bool {
	_, err := mysql.db.Exec(sql)
	if err != nil {
		return false
	}
	return true
}

//查询，返回map[string]string的格式
//怕注入使用fmt进行格式化后传入
func (mysql *Mysql) Select(dest interface{}, sql string) bool {
	err := mysql.db.Select(dest, sql)
	if err != nil {
		return false
	}
	return true
}
