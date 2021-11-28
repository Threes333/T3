package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB //数据库句柄
	err error   //错误信息
)

// @title 初始化数据库
// @description 初始化与数据库的连接与设置
func InitDB() {
	DB, err = sql.Open("mysql", "root:qazpl.123456@tcp(127.0.0.1:3306)/threes?charset=utf8")
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(100)
}
