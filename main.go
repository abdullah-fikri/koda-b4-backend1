package main

import (
	"backend1/controllers"
	"backend1/view"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(controllers.CorsMiddleware(r))
	r.Use(controllers.AllowPreflight(r))
	view.Routes(r)

	r.Run(":8081")
}
