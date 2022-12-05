package dao

import (
	"fmt"
	"message-board/model"
)

func InsertComment(t model.Comment) (err error) {
	_, err = DB.Exec("insert into comment(userName,pId,content)values(?,?,?)",
		t.UserName, t.PId, t.Content)
	if err != nil {
		fmt.Printf("When insert Comment,mysql Exec insert failed:%v ", err)
	}
	return err
}

func SelectComment(pId, userName string) (err error) {
	_, err = DB.Exec("select*from comment where FIND_IN_SET(id,getComment(?))", pId)
	if err != nil {
		fmt.Printf("When select Comment,mysql Exec insert failed:%v ", err)
	}
	return
}
func UpdateComment(t model.Comment) (err error) {
	stmt, err := DB.Prepare("update comment set content=? where userName=? and pId=?")
	if err != nil {
		fmt.Printf("When update ParentComment,mysql prepare failed:%v", err)
	}
	_, err = stmt.Exec(t.Content, t.UserName, t.PId)
	if err != nil {
		fmt.Printf("When update ParentComment,mysql query failed:%v", err)
	}
	return
}
