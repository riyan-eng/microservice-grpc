package datastruct

type ExampleList struct {
	Id        string `db:"uuid" json:"id"`
	Detail    string `db:"detail" json:"detail"`
	Name      string `db:"name" json:"name"`
	TotalRows int32  `db:"total_rows" json:"-"`
}

type ExampleDetail struct {
	Id     string `db:"uuid" json:"id"`
	Detail string `db:"detail" json:"detail"`
	Name   string `db:"name" json:"name"`
}
