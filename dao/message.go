package dao

import (
	"log"
	"message-board/model"
)

func InsertMessages(m model.Message) (err error) {
	//插入用户对应的留言
	DB.Exec("insert into message(username,)where(?,?,?,?) ",
		m.MessageId, m.AuthorId, m.ReceiveId, m.Detail)
	if err != nil {
		log.Println(err)
	}
	return
}
func SearchMessage(MessageId, AuthorId, ReceiveId int64) (m model.Message, err error) {
	//查询用户Id对应的留言
	row := DB.QueryRow("select detail from message where MessageId=? and AuthorId=? and ReceiveId=?",
		MessageId, AuthorId, ReceiveId)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&m.MessageId, &m.AuthorId, &m.ReceiveId, &m.Detail)
	return
}
func Update(m model.Message) (err error) {
	//更新用户对应的留言
	DB.Exec("update message set detail=? where MessageId=?and AuthorId=? and ReceiveId=? ",
		m.Detail, m.MessageId, m.AuthorId, m.ReceiveId)
	if err != nil {
		log.Println(err)
	}
	return
}
