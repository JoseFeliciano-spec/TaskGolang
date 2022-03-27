package database

import (
	"TaskGolang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn := "root@tcp(localhost)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error en la base de datos")
	}

	db.AutoMigrate(&models.User{})

	return db, err
}
