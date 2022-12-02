package dao

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"message-board/model"
	"time"
)

func InsertUser(u model.User) (err error) {
	//插入相关的用户信息及保密问题
	_, err = DB.Exec("insert into information(UserName,Password,Question,Answer)values(?,?,?,?) ",
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
func InsertPersonalInformation(u model.User) (err error) {
	_, err = DB.Exec("insert into personalInformation(userName,person)values(?,?)",
		u.UserName, u.Person)
	if err != nil {
		log.Printf("When insert personal information,mysql Exec insert failed:%v ", err)
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
	for row.Next() {
		err = row.Scan(&u.UserName, &u.Password, &u.Id, &u.Answer, &u.Question)
	}
	return
}
func SearchUserByQA(question, answer string) (u model.User, err error) {
	//预处理
	stmt, err := DB.Prepare("select * from information where Question=? and Answer=?")
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
	for row.Next() {
		err = row.Scan(&u.UserName, &u.Password, &u.Id, &u.Answer, &u.Question)

	}
	return
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
func GetJWT(password, userName string) (string, error) {
	mySigningkey := []byte(password)
	k := model.MyClaims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*24*7,
			Issuer:    userName,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, k)
	s, err := t.SignedString(mySigningkey)
	return s, err
}
