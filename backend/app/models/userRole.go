package model

import "time"

type UserRole struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	RoleID uint `json:"role_id" gorm:"index"` // Foreign key column with index
	Role   Role `gorm:"foreignkey:RoleID"`

	UserID uint `json:"user_id" gorm:"index"` // Foreign key column with index
	User   User `gorm:"foreignkey:UserID"`
}
