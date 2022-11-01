package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username    string   `json:"username" validate:"required"`
	Description string   `json:"descriptio"`
	Image       string   `json:"image"`
	UserId      uint     `json:"user_id"`
	Images      *[]Image `json:"images"`
}
