package panelAuthController

import (
	request "jora/app/http/requests"
	model "jora/app/models"
	"jora/database/postgres"
	"jora/utility"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {

	req := &request.PanelRegisterRequest{}
	if !request.Validation(c, req) {
		return
	}

	// check title unique in db
	var titleCount int64
	postgres.DB.Model(model.Company{}).Where("title = ?", req.Title).Count(&titleCount)
	if titleCount > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "title is duplicate!"})
		return
	}

	// check phone unique in db
	var phoneCount int64
	postgres.DB.Model(model.Company{}).Where("phone = ?", req.Phone).Count(&phoneCount)
	if phoneCount > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "phone is duplicate!"})
		return
	}

	newCompany := &model.Company{
		Title:    req.Title,
		Phone:    req.Phone,
		Password: req.Password,
	}

	newCompany.SetPassword()

	var token string
	var err error

	// Begin a transaction.
	postgres.DB.Transaction(func(tx *gorm.DB) error {
		// Check for transaction errors.
		if tx.Error != nil {
			panic(tx.Error)
		}

		// save company
		tx.Create(newCompany)

		token, err = utility.CreateToken(newCompany)
		if err != nil {

			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})

		}

		// return nil will commit the whole transaction
		return nil
	})
	// Commit the transaction if everything is successful.

	c.JSON(http.StatusOK, gin.H{"jwt_token": token})
}
