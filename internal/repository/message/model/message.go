package model

import (
	"database/sql"
)

type Message struct {
	Id        int64        `db:"id"`
	ChatId    int64        `db:"chat_id"`
	UserId    int64        `db:"user_id"`
	Text      string       `db:"text"`
	Timestamp sql.NullTime `db:"timestamp"`
}
