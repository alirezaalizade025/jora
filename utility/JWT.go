package utility

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)



type TokenDetails struct {
	gorm.Model

	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"access_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

func GenerateToken(user_id uint) (string, error) {
	var err error

	if CheckJwtTokenExists() != nil {
		return "", err
	}

	// // atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	// token, err := at.SignedString([]byte(Getenv("JWT_SECRET_KEY", "")))
	// if err != nil {
	// 	return "", err
	// }

	// return token, nil



	if err != nil {
		return "",err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func CheckJwtTokenExists() error {
	if Getenv("JWT_SECRET_KEY", "") == "" {
		log.Fatalf("SECRET not found in .env file")
		return errors.New("SECRET not found in .env file")
	}

	return nil
}

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// func DeleteAuth(givenUuid string) (int64, error) {
// 	deleted, err := client.Del(givenUuid).Result()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return deleted, nil
// }
