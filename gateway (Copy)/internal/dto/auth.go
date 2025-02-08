package dto

type AuthLogin struct {
	Username string `json:"username" valid:"required" minLength:"4" maxLength:"16" example:"admin" validate:"required"`
	Password string `json:"password" valid:"required" example:"12345678" validate:"required" enums:"foo,bar,baz"`
}

type AuthRefresh struct {
	Token string `json:"token" valid:"required"`
}
