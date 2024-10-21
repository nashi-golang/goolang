package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error
	dsn := "root:@tcp(localhost:3306)/wedding_diary?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Wedding{}, &People{})

	DB = db
}
