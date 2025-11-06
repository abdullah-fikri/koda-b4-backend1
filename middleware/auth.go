package middleware

import (
	"backend1/lib"
	"backend1/responses"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(401, responses.Response{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		payload, err := lib.VerifyToken(token)
		if err != nil {
			ctx.JSON(401, responses.Response{
				Success: false,
				Message: "Token invalid",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", payload)
		ctx.Next()

	}
}
