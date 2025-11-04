package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any
}

type Query struct {
	Name string `form:"name"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	var data Query
	// get
	var user []User
	r.GET("/users", func(ctx *gin.Context) {

		err := ctx.BindQuery(&data)
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Error",
			})
			return
		}

		ctx.JSON(200, Response{
			Success: true,
			Message: "Nama list",
			Data:    user,
		})
	})

	//get id
	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for _, u := range user {
			if u.Id == id {
				ctx.JSON(200, Response{
					Success: true,
					Message: "User ditemukan",
					Data:    u,
				})
				return
			}
		}
		ctx.JSON(404, Response{
			Success: false,
			Message: "User tidak ditemukan",
		})
	})
	//post
	r.POST("/users", func(ctx *gin.Context) {
		err := ctx.BindQuery(&data)
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Error",
			})
			return
		}
		id := ctx.PostForm("id")
		name := ctx.PostForm("name")

		user = append(user, User{Id: id, Name: name})

	})
	//delete
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i, u := range user {
			if u.Id == id {
				user = append(user[:i], user[i+1:]...)
				return
			}
		}
		ctx.JSON(404, Response{
			Success: false,
			Message: "tidak ditemukan id tsb",
		})
	})
	//patch
	r.PATCH("/users/:id", func(ctx *gin.Context) {
		idQ := ctx.PostForm("id")
		name := ctx.PostForm("name")

		id := ctx.Param("id")
		for i, u := range user {
			if u.Id == id {
				user[i] = User{Id: idQ, Name: name}
			}
		}

		ctx.JSON(404, Response{
			Success: false,
			Message: "tidak ditemukan id tsb",
		})
	})
	r.Run(":8081")
}
