package controllers

import (
	"backend1/models"
	"backend1/responses"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

var account []models.Accounts

func AuthController(r *gin.Engine) {
	r.POST("/auth/register", func(ctx *gin.Context) {
		idQ := ctx.PostForm("id")
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")

		argon := argon2.DefaultConfig()
		bytePassword, err := argon.HashEncoded([]byte(password))
		if err != nil {
			ctx.JSON(400, responses.Response{
				Success: false,
				Message: "Failed to hash password",
			})
			return
		}

		tmp := models.Accounts{
			Id:       idQ,
			Email:    email,
			Password: string(bytePassword),
		}

		account = append(account, tmp)

		ctx.JSON(200, responses.Response{
			Success: true,
			Message: "Register success",
		})
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
			if email == u.Email {
				ok, err := argon2.VerifyEncoded([]byte(password), []byte(u.Password))
				if err != nil || !ok {
					ctx.JSON(400, responses.Response{
						Success: false,
						Message: "wrong email or password",
					})
					return
				}

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
