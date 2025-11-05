package main

import (
	"backend1/controllers"
	"backend1/view"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "backend1/docs"
)

// @title API user & auth
// @version 1.0
// @description dokumentasi
// @host localhost:8081
// @BasePath /
func main() {
	r := gin.Default()
	r.Use(controllers.CorsMiddleware(r))
	r.Use(controllers.AllowPreflight(r))
	view.Routes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8081")
}
