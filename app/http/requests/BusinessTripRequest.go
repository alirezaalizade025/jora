package request

import (
	"github.com/gin-gonic/gin"
)

type BusinessTripRequest struct {
	FinishAt string `json:"finish_at" form:"finish_at" validate:"required,datetime=2006-01-02 15:04:05"`
	StartAt  string `json:"start_at" form:"start_at" validate:"required,datetime=2006-01-02 15:04:05"`

	Note string `json:"note" form:"note" validate:"required,max=255"`
}

func (r *BusinessTripRequest) bindValue(c *gin.Context) {

	// bind value from request to struct attribute for validation purpose
	var req BusinessTripRequest
	if err := c.ShouldBind(&req); err != nil {
		// handle error
		return
	}

	// select witch attribute to bind
	r.StartAt = req.StartAt
	r.FinishAt = req.FinishAt
	r.Note = req.Note
}
