package configs

import (
	"fmt"
	"os"

	"example.com/rest-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database")
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
