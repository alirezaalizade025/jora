package utility

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	db "jora/database/postgres"
)

type TokenDetails struct {
	gorm.Model

	UserID       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"access_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
	Revoke       bool   `json:"revoke" gorm:"default:false"`
}

func GenerateToken(user_id uint) (string, error) {
	var err error

	if CheckJwtTokenExists() != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["aud"] = generateUniqueUUID().String()
	claims["user_id"] = user_id

	// life time of token
	lifeTimeHours, _ := convertor.ToInt(Getenv("JWT_LIFE_TIME_HOURS", "1"))

	claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifeTimeHours)).Unix()

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

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)

	// validate token format
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func TokenCheckDb(c *gin.Context) error {
	tokenString := ExtractToken(c)

	claims := ExtractTokenClaim(tokenString)

	// expire on token data
	if claims["exp"] == nil || claims["user_id"] == "" {
		return errors.New("TOKEN IS INVALID")
	}

	expiredAt := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expiredAt) {
		return errors.New("TOKEN IS EXPIRED")
	}

	// expire on database (check uuid for that user id)
	var td TokenDetails
	row := db.DB.Where("access_token = ?", tokenString).Where("user_id = ?", claims["user_id"]).First(&td)

	// if token not exists
	if row.RowsAffected == 0 {
		return errors.New("TOKEN NOT FOUND")
	}

	// if token revoked
	if td.Revoke {
		return errors.New("TOKEN IS REVOKED")
	}

	// if token expired
	if time.Now().After(time.Unix(td.AtExpires, 0)) {
		return errors.New("TOKEN IS EXPIRED")
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

func ExtractTokenClaim(token string) jwt.MapClaims {

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		log.Panicln(err)
	}

	return claims

}

func generateUniqueUUID() uuid.UUID {
	uuidWithHyphen := uuid.New()

	for (db.DB.Where("access_uuid = ?", uuidWithHyphen.String()).First(&TokenDetails{}).RowsAffected != 0) {
		uuidWithHyphen = uuid.New()
	}

	return uuidWithHyphen
}

func Logout(access_token string) {
	var td TokenDetails
	db.DB.Where("access_token = ?", access_token).First(&td)

	td.Revoke = true

	db.DB.Save(&td)
}
