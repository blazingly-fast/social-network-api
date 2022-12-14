package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_name    string `json:"first_name" validate:"required"`
	Last_name     string `json:"last_name" validate:"required"`
	Password      string `json:"password" validate:"required"`
	Email         string `json:"email" validate:"email,required"`
	Phone         string `json:"phone" validate:"required"`
	Token         string `json:"token"`
	User_type     string `json:"user_type" validate:"required"`
	Refresh_token string `json:"refresh_token"`
	User_id       string `json:"user_id"`
	Account       *Account
}
