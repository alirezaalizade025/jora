package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userModel "nomasho/app/models/user"
	"nomasho/database/postgres"
	"nomasho/utility"
)

type User struct {
	RegisterNumber string   `json:"register_number" form:"register_number" binding:"required"`
	Password       string `json:"password" form:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var u User

	// db connection
	db := postgres.Connection()
	sqlDB, _ := db.Conn.DB()
	defer sqlDB.Close()

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	// find user by id in database
	user := userModel.User{}

	result := db.Conn.Where("register_number = ?", u.RegisterNumber).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}


	// compare the user from the request, with the one we defined:
	if user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token, error := utility.CreateToken(user.ID)
	if error != nil {
		c.JSON(http.StatusUnprocessableEntity, error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
