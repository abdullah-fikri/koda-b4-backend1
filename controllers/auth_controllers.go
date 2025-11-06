package controllers

import (
	"backend1/models"
	"backend1/responses"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
// @Summary Melihat semua akun yang sudah terdaftar (Pagination + Search)
// @Tags Auth
// @Produce json
// @Param search query string false "Cari email"
// @Param page query int false "Nomor halaman (default 1)"
// @Param limit query int false "Jumlah data per halaman (default 5)"
// @Success 200 {object} responses.Response{data=[]models.Accounts}
// @Router /auth/register [get]
func GetRegisteredUsers(ctx *gin.Context) {
	page := 1
	limit := 5
	search := ctx.Query("search")
	filtered := account
	if search != "" {
		filtered = []models.Accounts{}
		for _, x := range account {
			if strings.Contains(strings.ToLower(x.Email), strings.ToLower(search)) {
				filtered = append(filtered, x)
			}
		}
	}

	if ctx.Query("page") != "" {
		fmt.Sscan(ctx.Query("page"), &page)
	}
	if ctx.Query("limit") != "" {
		fmt.Sscan(ctx.Query("limit"), &limit)
	}
	start := (page - 1) * limit
	end := start + limit
	if start >= len(filtered) {
		ctx.JSON(200, responses.Response{
			Success: true,
			Message: "List akun",
			Data:    []models.Accounts{},
		})
		return
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	pagedData := filtered[start:end]

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

// UploadPicture godoc
// @Summary Update profile picture user
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "User ID"
// @Param picture formData file true "Profile picture (max 1MB)"
// @Success 200 {object} responses.Response{data=models.Accounts}
// @Router /auth/users/{id}/profile-picture [patch]
func UploadPicture(ctx *gin.Context) {
	id := ctx.Param("id")
	file, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(400, responses.Response{
			Success: false,
			Message: "failed",
		})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	formatExt := []string{".jpg", ".jpeg", ".png"}
	formatFile := false
	for _, x := range formatExt {
		if ext == x {
			formatFile = true
		}
	}
	if !formatFile {
		ctx.JSON(400, responses.Response{
			Success: false,
			Message: "hanya bisa upload gambar jpg, jpeg, dan png",
		})
		return
	}
	filename := "profile-picture-" + id + ".jpg"
	path := "./uploads/" + filename

	if file.Size > 1<<20 {
		ctx.JSON(400, responses.Response{
			Success: false,
			Message: "file to large",
		})
		return
	}
	ctx.SaveUploadedFile(file, path)
	for i, u := range account {
		if u.Id == id {
			account[i].ProfilePicture = path
			ctx.JSON(200, responses.Response{
				Success: true,
				Message: "Success save file",
				Data:    account[i],
			})
			return
		}
	}

	ctx.JSON(400, responses.Response{
		Success: false,
		Message: "id not found",
	})
}

func AuthController(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.POST("/register", Register)
	auth.GET("/register", GetRegisteredUsers)
	auth.POST("/login", Login)
	auth.PATCH("/users/:id/profile-picture", UploadPicture)
}
