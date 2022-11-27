package model

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
