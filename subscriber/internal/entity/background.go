package entity

type BackgroundCreate struct {
	Id   string               `json:"id"`
	Type string               `json:"type"`
	Body BackgroundCreateBody `json:"body"`
}

type BackgroundCreateBody struct {
	To          string `json:"to"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Var1        string `json:"v1,omitempty"`
	Var2        string `json:"v2,omitempty"`
	Var3        string `json:"v3,omitempty"`
	Var4        string `json:"v4,omitempty"`
	Var5        string `json:"v5,omitempty"`
}
