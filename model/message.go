package model

type Message struct {
	MessageId  int64  `json:"message-id"`
	Detail     string `json:"detail"`
	AuthorId   int64  `json:"authorId"`
	ReceiveId  int64  `json:"receiveId"`
	LikeNumber int64  `json:"likeNumber"`
}
