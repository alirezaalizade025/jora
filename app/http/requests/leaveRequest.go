package request

import (
	"github.com/gin-gonic/gin"
)

type LeaveRequest struct {
	ClockOut string `json:"clock_out" form:"clock_out" validate:"required,datetime=2006-01-02"`
	ClockIn  string `json:"clock_in" form:"clock_in" validate:"required,datetime=2006-01-02"`
	Type     string `json:"type" form:"type" validate:"required,oneof=sick_leave annual_leave vacation_leave"`

	Note string `json:"note" form:"note" validate:"omitempty,max=255"`
}

func (r *LeaveRequest) bindValue(c *gin.Context) {

	// bind value from request to struct attribute for validation purpose
	var req LeaveRequest
	if err := c.ShouldBind(&req); err != nil {
		// handle error
		return
	}

	// select witch attribute to bind
	r.ClockIn = req.ClockIn
	r.ClockOut = req.ClockOut
	r.Type = req.Type
	r.Note = req.Note
}
