package view

import (
	"backend1/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controllers.GetAllUsers(r)
	controllers.AuthController(r)
}
