package database

import (
	"MyGram/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "mygram.cygn9doqbcyr.ap-southeast-1.rds.amazonaws.com"
	user     = "postgres"
	password = "12345678"
	dbPort   = "5432"
	dbName   = "mygram"
	db       *gorm.DB
	err      error
)

// var (
// 	host     = os.Getenv("DB_HOST")
// 	user     = os.Getenv("DB_USER")
// 	password = os.Getenv("DB_PASSWORD")
// 	dbPort   = os.Getenv("DB_PORT")
// 	dbName   = os.Getenv("DB_NAME")
// 	db       *gorm.DB
// 	err      error
// )

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connection to database : ", err)
	}

	fmt.Println("Sukses Koneksi ke Database")
	db.Debug().AutoMigrate(model.User{}, model.Photo{}, model.Comment{}, model.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
