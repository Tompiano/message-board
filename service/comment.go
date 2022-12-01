package service

import (
	"message-board/dao"
	"message-board/model"
)

func CreateChildComment(t model.Comment) (err error) {
	err = dao.InsertChildComment(t)
	return err
}
func CreateParentComment(t model.Comment) (err error) {
	err = dao.InsertParentComment(t)
	return err
}
func SearchParentComment(MessageId, ParentUserId int64) (t model.Comment, err error) {
	_, err = dao.SelectParentComment(MessageId, ParentUserId)
	return
}
func SearchChildComment(MessageId, ParentId, ChildId int64) (t model.Comment, err error) {
	_, err = dao.SelectChildComment(MessageId, ParentId, ChildId)
	return
}
func ModifyParentComment(t model.Comment) (err error) {
	err = dao.UpdateParentComment(t)
	return err
}
func ModifyChildComment(t model.Comment) (err error) {
	err = dao.UpdateChildComment(t)
	return
}
