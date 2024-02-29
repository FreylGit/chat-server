package model

type User struct {
	Id       int64  `db:"id"`
	Username string `db:"username"`
}
