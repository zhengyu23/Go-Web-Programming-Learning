package data

import (
	"database/sql"
	"log"
)

// Db全局变量
var Db *sql.DB

// Web应用启动时对Db变量初始化的init函数
func init() {
	var err error
	Db, err := sql.Open("postgres", "dbname=chitchat sslomode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
