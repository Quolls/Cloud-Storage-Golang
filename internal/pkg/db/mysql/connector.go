package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	sql.Open("mysql", "root:root@tcp(127.0.0.1:3006/cloud_storage_sys?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}

func GetDb() *sql.DB {
	return db
}
