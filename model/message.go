package model

type Message struct {
	MessageId int64  `json:"message-id"`
	Detail    string `json:"detail"`
}
