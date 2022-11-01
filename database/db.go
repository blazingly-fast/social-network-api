package database

import (
	"log"

	"github.com/blazingly-fast/social-network-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Image{})

	return db
}
