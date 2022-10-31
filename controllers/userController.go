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
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User
		defer cancel()

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.WithContext(c).First(&foundUser, "email = ?", user.Email)

		if foundUser.ID == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
			return
		}

		passwordIsValid, msg := helpers.VerifyPassword(user.Password, foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.Email, foundUser.First_name, foundUser.Last_name, foundUser.User_type, foundUser.User_id)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		db.WithContext(c).Find(&foundUser)
		ctx.JSON(http.StatusOK, &foundUser)
	}
}

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		defer cancel()

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		password := helpers.HashPassword(user.Password)
		user.Password = password

		var count int64
		db.Model(&user).Where("email = ?", user.Email).Count(&count)
		defer cancel()
		if count > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exists"})
			return
		}

		token, refreshToken, _ := helpers.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.User_type, user.User_id)
		user.Token = token
		user.Refresh_token = refreshToken
		db.WithContext(c).Create(&user)
		ctx.JSON(http.StatusOK, &user.User_id)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// if err := helpers.CheckUserType(ctx, "ADMIN"); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
		// if err != nil || recordPerPage < 1 {
		// 	recordPerPage = 10
		// }
		// page, err1 := strconv.Atoi(ctx.Query("page"))
		// if err1 != nil || page < 1 {
		// 	page = 1
		// }

		// startIndex := (page - 1) * recordPerPage
		// startIndex, err = strconv.Atoi(ctx.Query("startIndex"))
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
