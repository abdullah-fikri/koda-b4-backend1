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

func main() {
	r := gin.Default()
	var data Query
	// get
	var names []string
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
			Data:    names,
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
		id := ctx.Query("id")
		name := ctx.Query("name")

		names = append(names, name, id)

	})
	//delete
	r.DELETE("/users", func(ctx *gin.Context) {})
	r.Run(":8081")
}
