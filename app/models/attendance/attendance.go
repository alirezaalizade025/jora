package attendance

import (
	"time"
)

type Attendance struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID uint `json:"user_id" gorm:"index"`

	StartAt  *time.Time `json:"start_at" gorm:"index"`
	FinishAt *time.Time `json:"finish_at" gorm:"index"`
	TypeInt  uint8      `json:"-" gorm:"index;type:SMALLINT CHECK (type >= 0);column:type"`
	Type     string     `json:"type" gorm:"-"`

	Note *string `json:"note" gorm:"type:text"`

	TeamLeadCheck *bool   `json:"team_lead_check" gorm:"default:null"`
	TeamLeadNote  *string `json:"team_lead_note" gorm:"type:text"`

	ManagerCheck *bool   `json:"manager_check" gorm:"default:null"`
	ManagerNote  *string `json:"manager_note" gorm:"type:text"`
}

func TYPE_MAP() map[int]string {
	return map[int]string{
		1: "working", // ساعت کاری

		11: "sick_leave",     // مرخصی استعلاجی
		12: "annual_leave",   // مرخصی استحقاقی-شخصی
		13: "vacation_leave", // مرخصی استحقاقی-تعطیلات

		14: "hourly_leave", // مرخصی ساعتی

		21: "business_trip", // ماموریت کاری
		22: "remote_work",   // دورکاری
	}
}

//==============================================================================//
//																			    //
//                                  Getters									    //
//																				//
//==============================================================================//

func (attendance *Attendance) GetType() string {
	return TYPE_MAP()[int(attendance.TypeInt)]
}

//==============================================================================//
//																			    //
//                                  Setters										//
//																			    //
//==============================================================================//

func (attendance *Attendance) SetType() {
	for key, value := range TYPE_MAP() {
		if value == attendance.Type {
			attendance.TypeInt = uint8(key)
		}
	}
}
