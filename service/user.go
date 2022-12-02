package service

import (
	"message-board/dao"
	"message-board/model"
)

func SearchUserByUserName(username, password string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(username, password)
	return
}
func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}
func CreatePersonInformation(u model.User) (err error) {
	err = dao.InsertPersonalInformation(u)
	return
}
func ModifyUser(u model.User) error {
	err := dao.InsertModifiedPassword(u)
	return err
}
func ForgetPassword(question, answer string) (u model.User, err error) {
	u, err = dao.SearchUserByQA(question, answer)
	return
}
func CompareHashPassword(password, hash string) bool {
	plainPwd := []byte(password)
	jud := dao.ComparePassword(hash, plainPwd)
	return jud
}
func GetJwt(password, userName string) (string, error) {
	s, err := dao.GetJWT(password, userName)
	return s, err
}
