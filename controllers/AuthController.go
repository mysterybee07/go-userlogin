package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/userlogin/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	// var userData models.Users

	err := c.BodyParser(&data)
	if err != nil {
		fmt.Println("Unable to parse")
	}

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

// func Login(c *fiber.Ctx) error {

// }