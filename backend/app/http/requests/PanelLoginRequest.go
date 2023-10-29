package request

import (
	"github.com/gin-gonic/gin"
)

type PanelLoginRequest struct {
	Phone 			string `json:"phone" form:"phone" validate:"required,len=11"`
	Password        string `json:"password" form:"password" validate:"required,min=4,max=20"`
}

func (r *PanelLoginRequest) bindValue(c *gin.Context) {

	// bind value from request to struct attribute for validation purpose
	var req PanelLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		// handle error
		return
	}

	// select witch attribute to bind
	r.Phone = req.Phone
	r.Password = req.Password
}
