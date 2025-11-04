package controllers

import (
	"backend1/models"
	"backend1/responses"

	"github.com/gin-gonic/gin"
)

var data models.Query
var user []models.User

func GetAllUsers(r *gin.Engine) {
	r.GET("/users", func(ctx *gin.Context) {

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
			Data:    user,
		})
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for _, u := range user {
			if u.Id == id {
				ctx.JSON(200, responses.Response{
					Success: true,
					Message: "User ditemukan",
					Data:    u,
				})
				return
			}
		}
		ctx.JSON(404, responses.Response{
			Success: false,
			Message: "User tidak ditemukan",
		})
	})

	r.POST("/users", func(ctx *gin.Context) {
		err := ctx.BindQuery(&data)
		if err != nil {
			ctx.JSON(400, responses.Response{
				Success: false,
				Message: "Error",
			})
			return
		}
		id := ctx.PostForm("id")
		name := ctx.PostForm("name")

		user = append(user, models.User{Id: id, Name: name})
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i, u := range user {
			if u.Id == id {
				user = append(user[:i], user[i+1:]...)
				return
			}
		}
		ctx.JSON(404, responses.Response{
			Success: false,
			Message: "tidak ditemukan id tsb",
		})
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		idQ := ctx.PostForm("id")
		name := ctx.PostForm("name")

		id := ctx.Param("id")
		for i, u := range user {
			if u.Id == id {
				user[i] = models.User{Id: idQ, Name: name}
			}
		}

		ctx.JSON(404, responses.Response{
			Success: false,
			Message: "tidak ditemukan id tsb",
		})
	})
}
