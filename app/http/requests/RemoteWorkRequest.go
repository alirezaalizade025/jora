package request

import (
	"github.com/gin-gonic/gin"
)

type RemoteWorkRequest struct {
	StartAt  string `json:"start_at" form:"start_at" validate:"required,datetime=2006-01-02 15:04:05"`
	FinishAt string `json:"finish_at" form:"finish_at" validate:"required,datetime=2006-01-02 15:04:05"`

	Note string `json:"note" form:"note" validate:"omitempty,max=255"`
}

func (r *RemoteWorkRequest) bindValue(c *gin.Context) {

	// bind value from request to struct attribute for validation purpose
	var req RemoteWorkRequest
	if err := c.ShouldBind(&req); err != nil {
		// handle error
		return
	}

	// select witch attribute to bind
	r.StartAt = req.StartAt
	r.FinishAt = req.FinishAt
	r.Note = req.Note
}
