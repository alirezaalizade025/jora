package utility

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
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

var (
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
)

func GetTokens() {

	PublicKey = loadPublicKey()

	PrivateKey = loadPrivateKey()
}

func loadPublicKey() *rsa.PublicKey {

	var pubBytes []byte
	var err error

	if  pubBytes = []byte(Getenv("PUBLIC_KEY", "")); len(pubBytes) == 0 {
		pubBytes, err = ioutil.ReadFile("public.pem")
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
	}


	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	return publicKey
}

func loadPrivateKey() *rsa.PrivateKey {

	var privBytes []byte
	var err error

	if  privBytes = []byte(Getenv("PRIVATE_KEY", "")); len(privBytes) == 0 {
		privBytes, err = ioutil.ReadFile("private.pem")
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
	}


	PrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	return PrivateKey
}

type TokenDetails struct {
	gorm.Model

	UserID       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"access_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
	Guard        string `json:"guard"`
	Revoke       bool   `json:"revoke" gorm:"default:false"`
}

func GenerateToken(user_id uint) (string, error) {
	var err error

	if PrivateKey != nil || PublicKey != nil {
		GetTokens()
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

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)


	return token.SignedString(PrivateKey)
}

func VerifyPassword(password, hashedPassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)

	// validate token format
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {

			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD: %v", token.Header["alg"])
		}

		return PublicKey, nil
	})

	if err != nil {
		return err
	}

	return nil
}

func TokenCheckDb(c *gin.Context) error {
	tokenString := ExtractToken(c)

	if tokenString == "" {
		return errors.New("TOKEN IS INVALID")
	}

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

	// set logged user id to context
	// todo: get user data from database and set to context
	c.Set("userId", claims["user_id"])

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
		return PublicKey, nil
	})


	if err != nil && err.Error() != "token has invalid claims: token is expired" { //todo: fix this
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
