package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Image_addr  string `json:"image_addr"`
	Description string `json:"description"`
	AccountId   uint   `json:"account_id"`
}
