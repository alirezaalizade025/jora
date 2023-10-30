package model

import (
	"jora/database/postgres"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	RegisterNumber uint `json:"register_number" gorm:"uniqueIndex"`

	CompanyID uint    `json:"company_id" gorm:"index"` // Foreign key column with index
	Company   Company `gorm:"foreignkey:CompanyID"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar" gorm:"default:null"`
	Password  string `json:"password"`

	TeamID uint `json:"team_id" gorm:"index"` // Foreign key column with index

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


func (u User) GetTeam() (team map[string]interface{}) {

	postgres.DB.Where("team_id = ?", u.TeamID).First(&team)
	return team
}
