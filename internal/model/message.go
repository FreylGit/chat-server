package model

import "time"

type Message struct {
	Id       int64
	ChatId   int64
	UserId   int64
	Text     string
	CreateAt time.Time
}
