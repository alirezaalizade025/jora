package user

import (
	"errors"
	"jora/database/postgres"

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

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := postgres.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}
