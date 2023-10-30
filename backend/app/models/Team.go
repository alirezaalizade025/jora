package model

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model

	RegisterNumber uint `json:"register_number" gorm:"uniqueIndex"`

	CompanyID uint    `json:"company_id" gorm:"index"` // Foreign key column with index
	Company   Company `gorm:"foreignkey:CompanyID"`

	Title string `json:"title"`

	FirstLeaderID uint `json:"first_leader_id" gorm:"index"` // Foreign key column with index
	FirstLeader   User `gorm:"foreignkey:FirstLeaderID"`

	SecondLeaderID uint `json:"second_leader_id" gorm:"index"` // Foreign key column with index
	SecondLeader   User `gorm:"foreignkey:SecondLeaderID"`

	Members []User `json:"members"`
}
