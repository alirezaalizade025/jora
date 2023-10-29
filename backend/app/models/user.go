package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	RegisterNumber string `json:"register_number" gorm:"uniqueIndex"`

	CompanyID uint    `json:"company_id" gorm:"index"` // Foreign key column with index
	Company   Company `gorm:"foreignkey:CompanyID"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar" gorm:"default:null"`
	Password  string `json:"password"`

	// each user can has many team leads
	TeamLeads  []*User `json:"team_leads" gorm:"many2many:team_leads"`
	TemMembers []*User `json:"team_members" gorm:"many2many:team_leads"`
}

// getGuard implements utility.authenticatable.
func (u User) GetGuard() string {
	return "api-user"
}

// getID implements utility.authenticatable.
func (u User) GetID() uint {
	return u.ID
}
