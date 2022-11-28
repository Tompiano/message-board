package service

import (
	"message-board/dao"
	"message-board/model"
)

func CreateMessage(m model.Message) error {
	err := dao.InsertMessages(m)
	return err
}
func SearchDetail(MessageId, AuthorId, ReceiveId int64) (m model.Message, err error) {
	m, err = dao.SearchMessage(MessageId, AuthorId, ReceiveId)
	return
}
