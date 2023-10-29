package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	request "jora/app/http/requests"
	userModel "jora/app/models/user"
	"jora/database/postgres"
	"jora/utility"
)

func Login(c *gin.Context) {
	
	req := &request.LoginRequest{}
	if !request.Validation(c , req) {
		return
	}

	token, err := LoginCheck(req.RegisterNumber, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func LoginCheck(registerNumber string, password string) (string, error) {

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

	if err != nil {
		return "", err
	}
	
	token, err := utility.GenerateToken(u.ID)
	
	if err != nil {
		return "", err
	}

	// todo: do action if with same client id and user id login again
	
	err = SaveUserLoginData(u.ID, token)

	return token, err
}

func SaveUserLoginData(user_id uint, tok string) error {

	db := postgres.DB

	if tok == "" {
		return errors.New("token is empty")
	}

	td := &utility.TokenDetails{}
	td.AccessToken = tok
	// extract expire time from token string
	claims := utility.ExtractTokenClaim(tok)
	td.AtExpires = int64(claims["exp"].(float64))

	td.UserID = user_id

	return db.Save(&td).Error
}