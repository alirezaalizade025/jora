package user

import (
	"errors"
	db "nomasho/database/postgres"

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


func (m *User) TableName() string {
    return "users"
}

func (u User) Find(id uint) (int, error) {
	result := db.Connection().Conn.Table(u.TableName()).Where("id = ?", id).First(&u)
	
	

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// handle record not found error
		return 404, errors.New("User not found")
	} else if result.Error != nil {
		// handle other errors
		return 500, result.Error
	} else {
		// record found
		return 200, nil
	}
}
