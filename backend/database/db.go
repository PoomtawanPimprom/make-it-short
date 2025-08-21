package database

import (
	"fmt"
	"github/PoomtawanPimprom/make-it-short/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = "5432"
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func ConnectDB() {
	//set up database connection here
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// migrate model
	db.AutoMigrate(&models.Url{})

	DB = db
}
