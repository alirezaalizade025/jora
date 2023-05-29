package utility

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userid uint) (string, error) {
	var err error


	if CheckJwtTokenExists() != nil {
		return "", err
	}


	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	// atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(Getenv("JWT_SECRET_KEY", "")))
	if err != nil {
	   return "", err
	}

	return token, nil
}

func CheckJwtTokenExists() error {
	if Getenv("JWT_SECRET_KEY", "") == "" {
		log.Fatalf("SECRET not found in .env file")
		return errors.New("SECRET not found in .env file")
	}

	return nil
}