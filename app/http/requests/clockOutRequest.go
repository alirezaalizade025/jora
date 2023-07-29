package request

import (
	"github.com/gin-gonic/gin"
)

type ClockOutRequest struct {
	ClockOut string `json:"clock_out" form:"clock_out" validate:"required,datetime=2006-01-02 15:04:05"`
}

func (r *ClockOutRequest) bindValue(c *gin.Context) {

	// bind value from request to struct attribute for validation purpose
	var req ClockOutRequest
	if err := c.ShouldBind(&req); err != nil {
		// handle error
		return
	}

	// select witch attribute to bind
	r.ClockOut = req.ClockOut
}


