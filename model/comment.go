package model

type Comment struct {
	Id           int64  `json:"Id"`
	ParentId     int64  `json:"parentId"`
	ParentUserId int64  `json:"ParentUserId"`
	ChildId      int64  `json:"childId"`
	MessageId    int64  `json:"MessageId"`
	UserName     string `json:"userName"`
	Content      string `json:"Content"`
}
