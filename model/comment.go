package model

type Comment struct {
	Id       string `json:"Id"`
	PId      string `json:"PId"`
	UserName string `json:"userName"`
	Content  string `json:"Content"`
}
