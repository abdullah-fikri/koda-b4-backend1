package controllers

import (
	"backend1/models"
	"backend1/responses"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matthewhartstonge/argon2"
)

func CorsMiddleware(r *gin.Engine) gin.HandlerFunc {
	godotenv.Load()
	env := os.Getenv("ORIGIN_URL")
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", env)
		ctx.Header("Access-Control-Allow-Methods", "GET,POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")
		ctx.Next()
	}
}

func AllowPreflight(r *gin.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
			return
		}
		ctx.Next()
	}
}

func hashPassword(password string) []byte {
	argon := argon2.DefaultConfig()
	bytePassword, _ := argon.HashEncoded([]byte(password))
	return bytePassword
}

var account []models.Accounts

// Register godoc
// @Summary Register akun baru
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body models.Accounts true "Register Data"
// @Success 200 {object} responses.Response{data=models.Accounts}
// @Router /auth/register [post]
func Register(ctx *gin.Context) {
	var req models.Accounts
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, responses.Response{
			Success: false,
			Message: "Invalid JSON",
		})
		return
	}

	req.Password = string(hashPassword(req.Password))
	account = append(account, req)

	ctx.JSON(200, responses.Response{
		Success: true,
		Message: "Register success",
		Data:    req,
	})
}

// GetRegisteredUsers godoc
// @Summary Melihat semua akun yang sudah terdaftar
// @Tags Auth
// @Produce json
// @Success 200 {object} responses.Response{data=[]models.Accounts}
// @Router /auth/register [get]
func GetRegisteredUsers(ctx *gin.Context) {
	ctx.JSON(200, responses.Response{
		Success: true,
		Message: "List akun",
		Data:    account,
	})
}

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body models.Accounts true "Login Data"
// @Success 200 {object} responses.Response{data=models.Accounts}
// @Router /auth/login [post]
func Login(ctx *gin.Context) {
	var req models.Accounts
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, responses.Response{
			Success: false,
			Message: "Invalid JSON",
		})
		return
	}

	for _, u := range account {
		ok, _ := argon2.VerifyEncoded([]byte(req.Password), []byte(u.Password))
		if u.Email == req.Email && ok {
			ctx.JSON(200, responses.Response{
				Success: true,
				Message: "Login sukses",
				Data:    u,
			})
			return
		}
	}

	ctx.JSON(404, responses.Response{
		Success: false,
		Message: "Wrong email or password",
	})
}

func AuthController(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.POST("/register", Register)
	auth.GET("/register", GetRegisteredUsers)
	auth.POST("/login", Login)
}
