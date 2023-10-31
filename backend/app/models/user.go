package model

import (
	"jora/database/postgres"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	RegisterNumber uint `json:"register_number" gorm:"uniqueIndex:idx_users_company_id"`
	
	CompanyID uint    `json:"company_id" gorm:"uniqueIndex:idx_users_company_id"`
	Company   Company `gorm:"foreignkey:CompanyID"`
	
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar" gorm:"default:null"`
	Password  string `json:"password"`

	TeamID uint `json:"team_id" gorm:"index"`
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
