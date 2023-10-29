package model

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	Title    string    `json:"title" gorm:"uniqueIndex"`
	Phone    string    `json:"phone"  gorm:"uniqueIndex"`
	CreditTo time.Time `json:"credit_to"`

	Password string `json:"password"`

	// each user can has many team leads
	Users []User `json:"users"`
}

// getGuard implements utility.authenticatable.
func (Company) GetGuard() string {
	return "admin-user"
}

// getID implements utility.authenticatable.
func (company Company) GetID() uint {
	return company.ID
}
