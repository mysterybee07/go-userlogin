package databases

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mysterybee07/userlogin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading.env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connected to database successfully")
	}
	DB = db
	db.AutoMigrate(
		&models.Users{},
	)
}
