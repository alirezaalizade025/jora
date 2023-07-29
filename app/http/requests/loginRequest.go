package request

import (
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	RegisterNumber string `json:"register_number" form:"register_number" validate:"required"`
	Password       string `json:"password" form:"password" validate:"required"`
}

func (r *LoginRequest) bindValue(c *gin.Context) {

	// bind value from request to struct attribute for validation purpose
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		// handle error
		return
	}

	// select witch attribute to bind
	r.RegisterNumber = req.RegisterNumber
	r.Password = req.Password
}
