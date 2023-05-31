package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model


	RegisterNumber string `json:"register_number" gorm:"uniqueIndex"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Avatar         string `json:"avatar" gorm:"default:null"`
	Password       string `json:"password"`
}
