package company

import (
	"time"
	"jora/app/models/user"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	Title     string `json:"title" gorm:"uniqueIndex"`
	CreditTo time.Time `json:"credit_to"`

	// each user can has many team leads
	Users  []*user.User `json:"users"`
}
