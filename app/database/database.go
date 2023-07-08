package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func SqlConnect() (db *sql.DB) {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_CONTAINER"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_NAME")) + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := sql.Open("mysql", connect)
	if err != nil {
		log.Fatal(err)
	}
	return
}
