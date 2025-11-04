package models

type Query struct {
	Name string `form:"name"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Accounts struct {
	Id       string `form:"name"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"min=8,max=20"`
}
