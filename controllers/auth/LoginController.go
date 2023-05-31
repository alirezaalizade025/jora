package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	userModel "nomasho/app/models/user"
	"nomasho/database/postgres"
	"nomasho/utility"
)

type LoginRequest struct {
	RegisterNumber string   `json:"register_number" form:"register_number" binding:"required"`
	Password       string `json:"password" form:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}


	token, err := LoginCheck(request.RegisterNumber, request.Password)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":token})
}



func LoginCheck(registerNumber string, password string) (string,error) {
	
	var err error

	u := userModel.User{}

	result := postgres.DB.Model(userModel.User{}).Where("register_number = ?", registerNumber).First(&u)


	if result.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	if result.Error != nil {
		return "", err
	}


	err = utility.VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token,err := utility.GenerateToken(u.ID)

	if err != nil {
		return "",err
	}

	return token,nil
}


