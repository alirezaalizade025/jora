package attendance

import (
	"time"
)

type Attendance struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID uint `json:"user_id" gorm:"index"`

	ClockIn  time.Time `json:"clock_in" gorm:"index"`
	ClockOut time.Time `json:"clock_out" gorm:"index"`
	TypeInt  uint8     `json:"type_int" gorm:"index;type:SMALLINT CHECK (type >= 0);column:type"`
	Type	 string    `json:"type" gorm:"-"`

	CheckInNote  string `json:"check_in_note" gorm:"type:text"`
	CheckOutNote string `json:"check_out_note" gorm:"type:text"`

	TeamLeadCheck     bool   `json:"team_lead_check" gorm:"default:false"`
	TeamLeadCheckNote string `json:"team_lead_check_note" gorm:"type:text"`

	ManagerCheck     bool   `json:"manager_check" gorm:"default:false"`
	ManagerCheckNote string `json:"manager_check_note" gorm:"type:text"`
}

func TYPE_MAP() map[int]string {
	return map[int]string{
		1: "working", // ساعت کاری

		2: "sick_leave", // مرخصی استعلاجی
		3: "personal_leave", // مرخصی استحقاقی-شخصی
		4: "business_leave", // مرخصی-کاری
		5: "vacation_leave", // مرخصی استحقاقی-تعطیلات

		6: "working_from_home",	// دورکاری
	}
}


//==============================================================================//
//																			    //
//                                  Getters									    //
//																				//
//==============================================================================//

func (attendance *Attendance) GetTypeName() string {
	return TYPE_MAP()[int(attendance.TypeInt)]
}

//==============================================================================//
//																			    //
//                                  Setters										//
//																			    //
//==============================================================================//

func (attendance *Attendance) SetTypeName() {
	for key, value := range TYPE_MAP() {
		if value == attendance.Type {
			attendance.TypeInt = uint8(key)
		}
	}
}

