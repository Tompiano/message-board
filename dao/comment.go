package dao

import (
	"fmt"
	"message-board/model"
)

func InsertComment(t model.Comment) (err error) {
	_, err = DB.Exec("insert into comment(parentId,parentUserId,childId,MessageId,userName,Content)values(?,?,?,?,?)",
		t.ParentId, t.ParentUserId, t.ChildId, t.MessageId, t.UserName, t.Content)
	if err != nil {
		fmt.Printf("mysql Exec insert failed:%v ", err)
	}
	return err
}
func SearchComment(t model.Comment) (err error) {
	
}
