package models

type Query struct {
	Name string `form:"name"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Accounts struct {
	Id       string `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"-" binding:"min=8,max=20"`
}
