package service

import (
	"message-board/dao"
	"message-board/model"
)

func CreateMessage(m model.Message) error {
	err := dao.InsertMessages(m)
	return err
}
func SearchDetail(MessageId int64) (m model.Message, err error) {
	m, err = dao.SearchMessage(MessageId)
	return
}
func UpdateMessage(m model.Message) error {
	err := dao.Update(m)
	return err
}
func DeleteMessage(m model.Message) error {
	err := dao.DeleteMessage(m)
	return err
}
