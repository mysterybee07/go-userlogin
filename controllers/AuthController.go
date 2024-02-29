package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/userlogin/databases"
	"github.com/mysterybee07/userlogin/models"
)

// var DB *gorm
func Register(c *fiber.Ctx) error {
	var data map[string]interface{}

	err := c.BodyParser(&data)
	if err != nil {
		fmt.Println("Unable to parse:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to parse request body",
		})
	}

	if len(data["password"].(string)) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password must be at least 6 characters long",
		})
	}

	user := models.Users{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     data["email"].(string),
		Password:  data["password"].(string), // Remember to hash the password before storing it in production
	}

	if err := databases.DB.Create(&user).Error; err != nil {
		log.Println("Failed to create user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":    user,
		"message": "Account registered successfully",
	})
}

// func Login(c *fiber.Ctx) error {
// 	var data map[string]interface{}
// 	if err := c.BodyParser(&data); err != nil {
// 		fmt.Println("Unable to parse")
// 	}

// 	var user *models.Users
// 	databases.DB.Where("email?", data["email"]).First(&user)
// 	if user.Id == 0 {
// 		c.Status(200)
// 		return c.JSON(fiber.Map{
// 			"message": "Account not found",
// 		})
// 	}

// 	if user.Password == data["password"].(string) {
// 		c.Status(200)
// 		return c.JSON(fiber.Map{
// 			"message": "Incorrect password",
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"message": "Login Successful",
// 	})

// }
