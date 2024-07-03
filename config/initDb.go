package config

import (
	entities "crud-go/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "user:password@tcp(localhost:3306)/go-db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	err = db.AutoMigrate(&entities.Tweet{})
	if err != nil {
		return nil, err
	}

	return db, nil
}