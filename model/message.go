package model

type Message struct {
	MessageId int64  `json:"message-id"`
	AuthorId  int64  `json:"author-id"`
	ReceiveId int64  `json:"receive-id"`
	Detail    string `json:"detail"`
}
