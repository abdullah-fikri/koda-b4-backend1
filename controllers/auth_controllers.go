package controllers

import (
	"backend1/models"
	"backend1/responses"
	"fmt"
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
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id formData string true "User ID"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {object} responses.Response{data=models.Accounts}
// @Router /auth/register [post]
func Register(ctx *gin.Context) {
	id := ctx.PostForm("id")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if email == "" || password == "" {
		ctx.JSON(400, responses.Response{
			Success: false,
			Message: "Email & Password wajib diisi",
		})
		return
	}

	req := models.Accounts{
		Id:       id,
		Email:    email,
		Password: string(hashPassword(password)),
	}

	account = append(account, req)

	ctx.JSON(200, responses.Response{
		Success: true,
		Message: "Register success",
		Data:    req,
	})
}

// GetRegisteredUsers godoc
// @Summary Melihat semua akun yang sudah terdaftar (Pagination)
// @Tags Auth
// @Produce json
// @Param page query int false "Nomor halaman (default 1)"
// @Param limit query int false "Jumlah data per halaman (default 5)"
// @Success 200 {object} responses.Response{data=[]models.Accounts}
// @Router /auth/register [get]
func GetRegisteredUsers(ctx *gin.Context) {
	page := 1
	limit := 5

	if ctx.Query("page") != "" {
		fmt.Sscan(ctx.Query("page"), &page)
	}
	if ctx.Query("limit") != "" {
		fmt.Sscan(ctx.Query("limit"), &limit)
	}
	start := (page - 1) * limit
	end := start + limit
	if start >= len(account) {
		ctx.JSON(200, responses.Response{
			Success: true,
			Message: "List akun",
			Data:    []models.Accounts{},
		})
		return
	}
	if end > len(account) {
		end = len(account)
	}

	pagedData := account[start:end]

	ctx.JSON(200, responses.Response{
		Success: true,
		Message: "List akun",
		Data:    pagedData,
	})
}

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {object} responses.Response{data=models.Accounts}
// @Router /auth/login [post]
func Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	for _, u := range account {
		ok, _ := argon2.VerifyEncoded([]byte(password), []byte(u.Password))
		if u.Email == email && ok {
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
