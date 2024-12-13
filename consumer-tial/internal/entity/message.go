package entity

type Message struct {
	Id     string      `json:"id"`
	Status string      `json:"-"`
	Type   string      `json:"type"`
	Body   MessageBody `json:"body"`
}

type MessageBody struct {
	To          string `json:"to"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Var1        string `json:"v1,omitempty"`
	Var2        string `json:"v2,omitempty"`
	Var3        string `json:"v3,omitempty"`
	Var4        string `json:"v4,omitempty"`
	Var5        string `json:"v5,omitempty"`
}
