package auth

import (
	resource "jora/app/http/resources/userInfo"
	models "jora/app/models"
	"jora/database/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)


func UserInfo(c *gin.Context) {

	authId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}


	user := models.User{}

	result := postgres.DB.Model(models.User{}).Where("id = ?", authId).First(&user)
	if result.Error == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "404 NOT FOUND!"})
	}


	c.JSON(http.StatusAccepted, resource.UserInfoResource(user))
}