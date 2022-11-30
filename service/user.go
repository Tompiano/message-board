package service

import (
	"message-board/dao"
	"message-board/model"
)

func SearchUserByUserName(username, Id string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(username, Id)
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
func SecretPassword(password string) string {
	NewPassword := dao.HashPassword(password)
	return NewPassword
}
func CompareHashPassword(password, hash string) bool {
	plainPwd := []byte(password)
	jud := dao.ComparePassword(hash, plainPwd)
	return jud
}
