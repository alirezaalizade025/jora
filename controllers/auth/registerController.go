package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var sampleSecretKey = []byte("sd1as3f32a1fa1f3s2af23saf1sa2SDAD")

func generateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodEdDSA)


	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	println(tokenString)
	return tokenString, nil

}

func Register(c *gin.Context) {
	token, err := generateJWT()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
