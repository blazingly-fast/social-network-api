package helpers

import (
	"log"
	"os"
	"time"

	"github.com/blazingly-fast/social-network-api/database"
	"github.com/blazingly-fast/social-network-api/models"
	"github.com/golang-jwt/jwt"
)

var db = database.Init(os.Getenv("DATABASE_URL"))

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (token string, refreshToken string, err error) {
	claims := SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err

}

func UpdateAllTokens(token string, refreshToken string, userId string) {
	var user models.User
	db.Where("user_id = ?", userId).First(&user)
	log.Print(&user)
	user.Token = token
	user.Refresh_token = refreshToken
	db.Save(&user)
}
