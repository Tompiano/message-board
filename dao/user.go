package dao

import (
	"message-board/model"
)

func InsertUser(u model.User) (err error) {
	//插入相关的用户信息及保密问题
	DB.Exec("insert into user(username,password,question,answer)where(?,?,?,?) ",
		u.UserName, u.Password, u.Question, u.Answer)
	return
}
func InsertModifiedPassword(u model.User) (err error) {
	//插入修改后的密码
	DB.Exec("update user set password=? where (?) ", u.Password, u.UserName)
	return
}
func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,username,password from user where username=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName, &u.Password)
	return
}
func SearchUserByQA(question, answer string) (u model.User, err error) {
	row := DB.QueryRow("select id,username,password from user where quesiton=? and answer=?", question, answer)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName)
	return
}
