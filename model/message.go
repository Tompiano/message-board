package model

type Message struct {
	messageId int    `json:"message-id"`
	authorId  int    `json:"author-id"`
	receiveId int    `json:"receive-id"`
	detail    string `json:"detail"`
}
