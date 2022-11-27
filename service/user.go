package service

import (
	"message-board/dao"
	"message-board/model"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(name)
	return
}
func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}
func ModifyUser(u model.User) error {
	err := dao.InsertModifiedPassword(u)
	return err
}
func ForgetPassword(question, answer string) (u model.User, err error) {
	u, err = dao.SearchUserByQA(question, answer)
	return
}
