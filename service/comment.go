package service

import (
	"message-board/dao"
	"message-board/model"
)

func CreateComment(t model.Comment) (err error) {
	err = dao.InsertComment(t)
	return err
}

func SearchComment(pId, userName string) (err error) {
	err = dao.SelectComment(pId, userName)
	return err
}

func ModifyComment(t model.Comment) (err error) {
	err = dao.UpdateComment(t)
	return err
}
