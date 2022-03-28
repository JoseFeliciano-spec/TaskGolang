package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	LastName string `json:"lastName" gorm:"not null"`
	Email    string `json:"email" gorm:"not null; uniqueIndex"`
	Password []byte `json:"password" gorm:"not null"`
	gorm.Model
}

type UserLogin struct {
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
