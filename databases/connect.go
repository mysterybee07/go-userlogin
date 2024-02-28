package databases

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mysterybee07/userlogin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global variable for database connection
var DB *gorm.DB

// Connect initializes the database connection and sets up migrations
func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database successfully")
	DB = db

	// AutoMigrate will automatically create the table based on the User struct
	db.AutoMigrate(
		&models.Users{},
	)
}
