package model

import "time"

type Message struct {
	Id       int64     `db:"id"`
	UserId   int64     `db:"user_id"`
	ChatId   int64     `db:"chat_id"`
	Text     string    `db:"text"`
	CreateAt time.Time `db:"create_at"`
}
