package dto

type UserRegDto struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
