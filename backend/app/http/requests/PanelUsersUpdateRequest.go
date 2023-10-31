package request

import (
	"github.com/gin-gonic/gin"
)

type PanelUsersUpdateRequest struct {
	FirstName *string `json:"first_name" form:"first_name" validate:"omitempty,max=50"`
	LastName  *string  `json:"last_name" form:"last_name" validate:"omitempty,max=50"`

	RegisterNumber *uint `json:"register_number" form:"register_number" validate:"omitempty,min=0,max=15"`
	TeamId         *uint `json:"team_id" form:"team_id" validate:"omitempty"`
}

func (r *PanelUsersUpdateRequest) bindValue(c *gin.Context) {

	// bind value from request to struct attribute for validation purpose
	var req PanelUsersUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		// handle error
		return
	}

	// select witch attribute to bind
	r.FirstName = req.FirstName
	r.LastName = req.LastName
	r.RegisterNumber = req.RegisterNumber
	r.TeamId = req.TeamId
}
