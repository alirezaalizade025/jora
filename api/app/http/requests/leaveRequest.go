package request

import (
	"github.com/gin-gonic/gin"
)

type LeaveRequest struct {
	FinishAt string `json:"finish_at" form:"finish_at" validate:"required,datetime=2006-01-02"`
	StartAt  string `json:"start_at" form:"start_at" validate:"required,datetime=2006-01-02"`
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
	r.StartAt = req.StartAt
	r.FinishAt = req.FinishAt
	r.Type = req.Type
	r.Note = req.Note
}
