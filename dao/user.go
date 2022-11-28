package dao

import (
	"message-board/model"
	"message-board/util"
)

func InsertUser(u model.User) (err error) {
	//插入相关的用户信息及保密问题
	result, err := DB.Exec("insert into information(username,password,question,answer)where(?,?,?,?) ",
		u.UserName, u.Password, u.Question, u.Answer)
	util.Err(err)
	// 返回新插入数据的id
	result.LastInsertId()
	// 返回影响的行数
	result.RowsAffected()
	return
}
func InsertModifiedPassword(u model.User) (err error) {
	//插入修改后的密码
	result, err := DB.Exec("update information set password=? where (?) ", u.Password, u.UserName)
	util.Err(err)
	// 返回新插入数据的id
	result.LastInsertId()
	// 返回影响的行数
	result.RowsAffected()
	return
}
func SearchUserByUserName(username string) (u model.User, err error) {
	//预处理
	stmt, err := DB.Prepare("select id,username,password from information where username=?")
	util.Err(err)
	row, err := stmt.Query(username)
	util.Err(err)
	defer row.Close() //延迟关闭
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName, &u.Password)
	return
}
func SearchUserByQA(question, answer string) (u model.User, err error) {
	//预处理
	stmt, err := DB.Prepare("select id,username,password from information where quesiton=? and answer=?")
	util.Err(err)
	row, err := stmt.Query(question, answer)
	util.Err(err)
	defer row.Close() //延迟关闭
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName)
	return
}
