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

	dsn := "host=localhost user=postgres password=123456789 dbname=fiverrv2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")

	}

	DBConn = db
	DBConn.AutoMigrate(&models.Category{}, &models.Gig{}, &models.Comment{}, &models.Tags{}, &models.User{})

	fmt.Println("Connection Opened to Database")

}
