package model

type ChatUser struct {
	ChatId int64 `db:"chat_id"`
	UserId int64 `db:"user_id"`
}
