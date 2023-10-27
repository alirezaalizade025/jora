package teamLead

import (
	"time"
)

type TeamLead struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

    UserID     uint `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID"`
    TeamLeadID uint `gorm:"constraint:OnDelete:CASCADE;foreignKey:TeamLeadID"`
}
