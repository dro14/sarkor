package database

import (
	"log"

	"github.com/dro14/sarkor/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	var err error
	DB, err = gorm.Open(sqlite.Open("sarkor.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Phone{})
	if err != nil {
		log.Fatal(err)
	}
}
