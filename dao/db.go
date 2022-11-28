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
	err = db.Ping() //检查数据库是否可用且可访问
	if err != nil {
		fmt.Println("数据库连接失败")
		log.Println(err)
	}
}
