package postgres

import (
	"fmt"
	"net/http"
	"strconv"

	"nomasho/utility"
	userModel "nomasho/app/models/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
)



func CreateDatabase(databaseName string) {
	host := utility.Getenv("DB_HOST", "127.0.0.1")
	port := utility.Getenv("DB_PORT", "5432")
	user := utility.Getenv("DB_USER", "postgres")
	password := utility.Getenv("DB_PASSWORD", "postgres")
	timezone := utility.Getenv("DB_TIMEZONE", "Asia/Tehran")

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

func ConnectDataBase(){	

	CreateDatabase(utility.Getenv("DB_NAME", "nomasho_db"))
	
	host := utility.Getenv("DB_HOST", "127.0.0.1")
	port := utility.Getenv("DB_PORT", "5432")
	dbname := utility.Getenv("DB_NAME", "nomasho_db")
	user := utility.Getenv("DB_USER", "postgres")
	password := utility.Getenv("DB_PASSWORD", "postgres")
	sslmode := utility.Getenv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", host, port, dbname, user, password, sslmode)
	

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
		PrepareStmt: true,
	})

	CheckError(err)

	DB = db
	

	// migrations
	DB.AutoMigrate(&userModel.User{})	
}



func TestConnection() *gorm.DB {
	host := utility.Getenv("DB_HOST", "127.0.0.1")
	port := utility.Getenv("DB_PORT", "5432")
	user := utility.Getenv("DB_USER", "postgres")
	password := utility.Getenv("DB_PASSWORD", "postgres")
	sslmode := utility.Getenv("DB_SSLMODE", "disable")
	dbname := utility.Getenv("DB_NAME", "nomasho_test_db")

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



