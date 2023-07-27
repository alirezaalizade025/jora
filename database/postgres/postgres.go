package postgres

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabase(databaseName string) {
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "5432")
	user := getenv("DB_USER", "postgres")
	password := getenv("DB_PASSWORD", "postgres")
	timezone := getenv("DB_TIMEZONE", "Asia/Tehran")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s", host, port, user, password, timezone)
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// createDatabaseCommand := fmt.Sprintf("DROP DATABASE %s", databaseName)
	// DB.Exec(createDatabaseCommand)

	var databaseExists int64
	DB.Raw("SELECT count(*) FROM pg_database where datname = ?", databaseName).Count(&databaseExists)

	if databaseExists == 0 {
		createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", databaseName)
		DB.Exec(createDatabaseCommand)
	}
}

var DB *gorm.DB

func ConnectDataBase() {

	CreateDatabase(getenv("DB_NAME", "jora_db"))

	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "5432")
	dbname := getenv("DB_NAME", "jora_db")
	user := getenv("DB_USER", "postgres")
	password := getenv("DB_PASSWORD", "postgres")
	sslmode := getenv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", host, port, dbname, user, password, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
		PrepareStmt: true,
	})

	CheckError(err)

	DB = db
}

func TestConnection() *gorm.DB {
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "5432")
	user := getenv("DB_USER", "postgres")
	password := getenv("DB_PASSWORD", "postgres")
	sslmode := getenv("DB_SSLMODE", "disable")
	dbname := getenv("DB_NAME", "jora_test_db")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})

	CheckError(err)

	return db
}

func CheckError(err error) {
	if err != nil {
		// sentry.CaptureException(errors.New(err.Error()))
		// sentry.Flush(time.Second * 10)
		panic(err)
	}
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("per_page"))
		switch {
		case pageSize > 30:
			pageSize = 30
		case pageSize <= 0:
			pageSize = 30
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// to import cycle force copy here
func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
