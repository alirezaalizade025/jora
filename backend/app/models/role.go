package model

import "time"

type Role struct {

	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	CompanyID uint    `json:"company_id" gorm:"index"` // Foreign key column with index
	Company   Company `gorm:"foreignkey:CompanyID"`

	Title string `json:"title"`

	Users []User `gorm:"many2many:user_roles;"`
}
