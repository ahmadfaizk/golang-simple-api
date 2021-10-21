package db

import (
	"fmt"
	"golang-simple-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=go_simple_api port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB error: ", err)
	}
	return db
}

func AutoMigrate(db *gorm.DB)  {
	err := db.AutoMigrate(
		&model.Article{})
	if err != nil {
		return
	}
}