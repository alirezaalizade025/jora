package panelAuthController

import (
	resource "jora/app/http/resources/panel"
	models "jora/app/models"
	"jora/database/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminInfo(c *gin.Context) {

	authId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}


	company := models.Company{}

	result := postgres.DB.Model(models.Company{}).Where("id = ?", authId).First(&company)
	if result.Error == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "404 NOT FOUND!"})
	}


	c.JSON(http.StatusAccepted, resource.CompanyInfoResource(company))
}
