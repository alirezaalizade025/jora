package utility

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func LoadDotEnv() {
	err := godotenv.Load()

	if err != nil && err.Error() != "open .env: no such file or directory" {  //todo: refactor
		log.Fatalf("Some error occurred. Err: %s", err)
	}
}


func Getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func DBResponseHandle(result *gorm.DB) (int, error)  {
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// handle record not found error
		return 404, errors.New("User not found")
	} else if result.Error != nil {
		// handle other errors
		return 500, result.Error
	} else {
		// record found
		return 200, nil
	}
}

func InArray(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
