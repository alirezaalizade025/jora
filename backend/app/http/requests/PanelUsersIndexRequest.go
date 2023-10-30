package request

import (
	"github.com/gin-gonic/gin"
)

type PanelUsersIndexRequest struct {
	User   string `json:"user" form:"user" validate:"omitempty,max=50"`
	TeamId *uint  `json:"team_id" form:"team_id" validate:"omitempty"`

	Page    int `json:"page" form:"page" validate:"omitempty,min=0"`
	PerPage int `json:"per_page" form:"per_page" validate:"omitempty,min=1,max=30"`
}

func (r *PanelUsersIndexRequest) bindValue(c *gin.Context) {

	// bind value from request to struct attribute for validation purpose
	var req PanelUsersIndexRequest
	if err := c.ShouldBind(&req); err != nil {
		// handle error
		return
	}

	// select witch attribute to bind
	r.User = req.User
	r.TeamId = req.TeamId
	r.Page = req.Page
	r.PerPage = req.PerPage
}
