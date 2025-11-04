package controllers

import (
	"backend1/models"
	"backend1/responses"

	"github.com/gin-gonic/gin"
)

var account []models.Accounts

func AuthController(r *gin.Engine) {
	r.POST("/auth/register", func(ctx *gin.Context) {
		idQ := ctx.PostForm("id")
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")
		account = append(account, models.Accounts{Id: idQ, Email: email, Password: password})
	})

	r.GET("/auth/register", func(ctx *gin.Context) {

		err := ctx.BindQuery(&data)
		if err != nil {
			ctx.JSON(400, responses.Response{
				Success: false,
				Message: "Error",
			})
			return
		}

		ctx.JSON(200, responses.Response{
			Success: true,
			Message: "Nama list",
			Data:    account,
		})
	})

	r.POST("/auth/login", func(ctx *gin.Context) {
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")
		for _, u := range account {
			if u.Email == email && u.Password == password {
				ctx.JSON(200, responses.Response{
					Success: true,
					Message: "login sukses",
					Data:    u,
				})
				return
			}
		}
		ctx.JSON(404, responses.Response{
			Success: false,
			Message: "wrong email or password",
		})
	})
}
