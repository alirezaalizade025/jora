package auth

// import (
// 	"net/http"
// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/gin-gonic/gin"
// 	"nomasho/config"
// 	"nomasho/models"
// 	"nomasho/utils"
// )

// // LoginController is the controller for the login route
// func LoginController(c *gin.Context) {
// 	var login models.Login
// 	if err := c.ShouldBindJSON(&login); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// check if user exists
// 	var user models.User
// 	config.DB.Where("email = ?", login.Email).First(&user)
// 	if user.ID == 0 {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}

// 	// check if password is correct
// 	if err := utils.VerifyPassword(user.Password, login.Password); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}

// 	// create jwt
// 	token, err := utils.CreateToken(user.ID)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error creating token"})
// 		return
// 	}

// 	// send response
// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }
