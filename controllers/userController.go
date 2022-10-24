package controllers

import (
	"context"
	"net/http"
	"os"

	"time"

	"github.com/blazingly-fast/social-network-api/database"
	"github.com/blazingly-fast/social-network-api/helpers"
	"github.com/blazingly-fast/social-network-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var db = database.Init(os.Getenv("DATABASE_URL"))
var validate = validator.New()

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		defer cancel()

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		var count int64
		db.Model(&user).Where("email = ?", user.Email).Count(&count)
		defer cancel()

		if count > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exists"})
		}

		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken
		db.WithContext(c).Create(&user)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		if err := helpers.MatchUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		defer cancel()

		db.WithContext(c).Find(&user)
		ctx.JSON(http.StatusOK, &user)
	}
}
