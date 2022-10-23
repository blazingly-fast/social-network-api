package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func HashPassword(ctx *gin.Context, password string) error {
	return nil
}

func CheckPassword(ctx *gin.Context) error {
	return nil
}

func CheckUserType(ctx *gin.Context, role string) (err error) {
	userType := ctx.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(ctx *gin.Context, userId string) (err error) {
	userType := ctx.GetString("user_type")
	uid := ctx.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(ctx, userType)
	return err
}
