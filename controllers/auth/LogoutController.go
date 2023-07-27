package auth

import (
	"jora/utility"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	tokenString := utility.ExtractToken(c)

	utility.Logout(tokenString)

	c.JSON(200, gin.H{})
}
