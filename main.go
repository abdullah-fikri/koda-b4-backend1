package main

import (
	"backend1/view"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	view.Routes(r)

	r.Run(":8081")
}
