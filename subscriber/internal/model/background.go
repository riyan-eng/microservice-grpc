package model

type Background struct {
	Id     string `db:"id"`
	Status string `db:"status"`
	Type   string `db:"type"`
	Body   string `db:"body"`
	Error  string `db:"error"`
}
