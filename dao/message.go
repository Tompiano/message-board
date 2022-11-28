package dao

import (
	"log"
	"message-board/model"
)

func InsertMessages(m model.Message) (err error) {
	//插入用户对应的留言
	DB.Exec("insert into message(username,)where(?,?) ",
		m.MessageId, m.Detail)
	if err != nil {
		log.Println(err)
	}
	return
}
func SearchMessage(MessageId int64) (m model.Message, err error) {
	//查询用户Id对应的留言
	row := DB.QueryRow("select detail from message where MessageId=? ",
		MessageId)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&m.MessageId, &m.Detail)
	return
}
func Update(m model.Message) (err error) {
	//更新用户对应的留言
	DB.Exec("update message set detail=? where MessageId=? ",
		m.Detail, m.MessageId)
	if err != nil {
		log.Println(err)
	}
	return
}
func DeleteMessage(m model.Message) (err error) {
	//删除留言，并增加上”该留言已删除“字样
	DB.Exec("update message set detail=? where MessageId=?",
		m.Detail, m.MessageId)
	if err != nil {
		log.Println(err)
	}
	return
}
