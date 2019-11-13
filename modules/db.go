package modules

import (
	"database/sql"
	"github.com/carl-xiao/short-link-go/setting"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

/**
初始化数据库
*/
func Dbinit() *sql.DB {
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbName := sec.Key("NAME").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	url := user + ":" + password + "@tcp" + "(" + host + ")" + "/" + dbName + "?parseTime=true"
	dbUrl, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}
	//增加配置文件
	dbUrl.SetMaxOpenConns(100)
	dbUrl.SetMaxIdleConns(20)
	dbUrl.SetConnMaxLifetime(time.Second * 30)

	err = dbUrl.Ping()
	if err != nil {
		log.Fatal("mysql not connect")
	}
	return dbUrl
}
