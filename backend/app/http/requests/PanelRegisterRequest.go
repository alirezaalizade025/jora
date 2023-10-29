package request

import (
	"github.com/gin-gonic/gin"
)

type PanelRegisterRequest struct {
	Title           string `json:"title" form:"title" validate:"required,min=2,max=50"`
	Phone 			string `json:"phone" form:"phone" validate:"required,len=11"`

	Password        string `json:"password" form:"password" validate:"required,min=4,max=20"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password"`
}

func (r *PanelRegisterRequest) bindValue(c *gin.Context) {

	// bind value from request to struct attribute for validation purpose
	var req PanelRegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		// handle error
		return
	}

	// select witch attribute to bind
	r.Title = req.Title
	r.Phone = req.Phone
	r.Password = req.Password
	r.ConfirmPassword = req.ConfirmPassword
}
