package service

import (
	"message-board/dao"
	"message-board/model"
)

func CreateComment(t model.Comment) (err error) {
	err = dao.InsertComment(t)
	return err
}
