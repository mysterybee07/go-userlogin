package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mysterybee07/userlogin/databases"
	"github.com/mysterybee07/userlogin/routes"
)

func main() {
	databases.Connect()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading.env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
