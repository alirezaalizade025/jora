package panelAuthController

import (
	"errors"
	request "jora/app/http/requests"
	model "jora/app/models"
	"jora/database/postgres"
	"jora/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	req := &request.PanelLoginRequest{}
	if !request.Validation(c, req) {
		return
	}


	token, err := LoginCheck(req.Phone, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone or password is incorrect."})
		return
	}


	c.JSON(http.StatusOK, gin.H{"jwt_token": token})
}

func LoginCheck(mobile string, password string) (string, error) {

	var err error

	company := model.Company{}

	result := postgres.DB.Model(model.Company{}).Where("phone = ?", mobile).First(&company)
	if result.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	if result.Error != nil {
		return "", err
	}


	err = utility.VerifyPassword(password, company.Password)

	if err != nil {
		return "", err
	}
	
	// todo: do action if with same client id and user id login again
	return utility.CreateToken(company)
}