package dao

import (
	"fmt"
	"message-board/model"
)

func InsertParentComment(t model.Comment) (err error) {
	_, err = DB.Exec("insert into ParentComment(ParentId,ParentUserId,MessageId,UserName,Content)values(?,?,?,?,?)",
		t.ParentId, t.ParentUserId, t.MessageId, t.UserName, t.Content)
	if err != nil {
		fmt.Printf("When insert Parent Comment,mysql Exec insert failed:%v ", err)
	}
	return err
}
func InsertChildComment(t model.Comment) (err error) {
	_, err = DB.Exec("insert into ChildComment(ParentId,ChildId,MessageId,UserName,Content)values(?,?,?,?,?)",
		t.ParentId, t.ChildId, t.MessageId, t.UserName, t.Content)
	if err != nil {
		fmt.Printf("When insert Child Comment,mysql Exec insert failed:%v ", err)
	}
	return err
}
func SelectParentComment(MessageId, ParentUserId int64) (t model.Comment, err error) {
	stmt, err := DB.Prepare("select Content from ParentComment where MessageId=? and ParentUserId=?")
	if err != nil {
		fmt.Printf("When select ParentComment,mysql prepare failed:%v", err)
	}
	row, err := stmt.Query(MessageId, ParentUserId)
	if err != nil {
		fmt.Printf("When select ParentComment,mysql query failed:%v", err)
	}
	defer row.Close() //延迟关闭
	for row.Next() {
		err = row.Scan(&t.Content, &t.MessageId, &t.ParentUserId)
	}
	return
}
func SelectChildComment(MessageId, ParentId, ChildId int64) (t model.Comment, err error) {
	stmt, err := DB.Prepare("select Content from ChildComment where MessageId=?and ParentId=? and ChildId=?")
	if err != nil {
		fmt.Printf("When select ChildComment,mysql prepare failed:%v", err)
	}
	row, err := stmt.Query(MessageId, ParentId, ChildId)
	if err != nil {
		fmt.Printf("When select ChildComment,mysql query failed:%v", err)
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&t.Content, &t.MessageId, &t.ParentId, &t.ChildId)
	}
	return
}
func UpdateParentComment(t model.Comment) (err error) {
	stmt, err := DB.Prepare("update ParentComment set content=? where MessageId=? and ParentUserId=?")
	if err != nil {
		fmt.Printf("When update ParentComment,mysql prepare failed:%v", err)
	}
	_, err = stmt.Exec(t.Content, t.MessageId, t.ParentUserId)
	if err != nil {
		fmt.Printf("When update ParentComment,mysql query failed:%v", err)
	}
	return
}
func UpdateChildComment(t model.Comment) (err error) {
	stmt, err := DB.Prepare("update ChildComment set content=? where MessageId=? and ParentId=? and ChildId=?")
	if err != nil {
		fmt.Printf("When update ChildComment,mysql prepare failed:%v", err)
	}
	_, err = stmt.Exec(t.Content, t.MessageId, t.ParentId, t.ChildId)
	if err != nil {
		fmt.Printf("When update ChildComment,mysql query failed:%v", err)
	}
	return
}
