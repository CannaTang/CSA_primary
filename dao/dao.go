package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dB *sql.DB

const (
	user     = "root"
	password = "root"
	host     = "localhost"
	port     = 3306
	dbName   = "Q-A"
)

//链接上数据库
func init() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName))
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)
	//检查数据库是否可用可访问
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	dB = db
}
