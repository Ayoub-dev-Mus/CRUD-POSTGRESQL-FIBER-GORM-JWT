package database

import (
	"FIVERR/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {

	dsn := "host=ec2-3-93-206-109.compute-1.amazonaws.com user=ndrqvkdtymlxfi password=3c5452091d7680703b3cac6239820050914519b417c2af6f22a1cc5919cbea13 dbname=d423df9qigjj8s port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")

	}

	DBConn = db
	DBConn.AutoMigrate(&models.Category{}, &models.Gig{}, &models.Comment{}, &models.Tags{}, &models.User{}, &models.Order{})

	fmt.Println("Connection Opened to Database")

}
