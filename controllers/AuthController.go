package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/userlogin/databases"
	"github.com/mysterybee07/userlogin/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	// var userData models.Users

	err := c.BodyParser(&data)
	if err != nil {
		fmt.Println("Unable to parse")
	}

	if len(data["password"].(string)) < 6 {
		c.Status(200)
		return c.JSON(fiber.Map{
			"message": "Password too short",
		})
	}

	// first_name, err := data["first_name"].(string)

	user := models.Users{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     data["email"].(string),
		Password:  data["password"].(string),
	}
	c.Status(201)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account registerd successfully",
	})

}

func Login(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse")
	}

	var user *models.Users
	databases.DB.Where("email?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(200)
		return c.JSON(fiber.Map{
			"message": "Account not found",
		})
	}

	if user.Password == data["password"].(string) {
		c.Status(200)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Login Successful",
	})

}
