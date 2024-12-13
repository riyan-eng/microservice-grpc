package model

type User struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Username string `db:"username"`
}
