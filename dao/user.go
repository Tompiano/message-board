package dao

import (
	"message-board/model"
)

func InsertUser(u model.User) (err error) {
	//插入相关的用户信息及保密问题
	DB.Exec("insert into user(name,password,question,answer)where(?,?,?,?) ",
		u.UserName, u.Password, u.Question, u.Answer)
	return
}
func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,name,password from user where name=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName, &u.Password)
	return
}
