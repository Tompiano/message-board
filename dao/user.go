package dao

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"message-board/model"
)

func InsertUser(u model.User) (err error) {
	//插入相关的用户信息及保密问题
	_, err = DB.Exec("insert into information(UserName,Password,Question,Answer)values(?,?,?) ",
		u.UserName, u.Password, u.Question, u.Answer)
	if err != nil {
		fmt.Printf("mysql Exec insert failed:%v ", err)
	}
	return
}
func InsertModifiedPassword(u model.User) (err error) {
	//插入修改后的密码
	_, err = DB.Exec("update information set password=? where (?) ", u.Password, u.UserName)
	if err != nil {
		fmt.Printf("mysql Exec update failed:%v ", err)
	}

	return
}
func SearchUserByUserName(username, password string) (u model.User, err error) {
	//预处理
	stmt, err := DB.Prepare("select * from information where UserName=? and Password=? ")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
	}
	row, err := stmt.Query(username, password)
	if err != nil {
		fmt.Printf("mysql query failed:%v", err)
	}
	defer row.Close() //延迟关闭
	if err = row.Err(); row.Err() != nil {
		return
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&u.UserName, &u.Password)
	}
	return
}
func SearchUserByQA(question, answer string) (u model.User, err error) {
	//预处理
	stmt, err := DB.Prepare("select Id,UserName,Password from information where Quesiton=? and Answer=?")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
	}
	row, err := stmt.Query(question, answer)
	if err != nil {
		fmt.Printf("mysql query failed:%v", err)
	}
	defer row.Close() //延迟关闭
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName)
	return
}
func HashPassword(Password string) string {
	password := []byte(Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.MaxCost)
	return string(hashedPassword)
}
func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
