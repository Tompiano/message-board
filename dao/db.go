package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Fatalf("connet mysql error :%v", err)
	}
	DB = db
	fmt.Println(db.Ping())
}
