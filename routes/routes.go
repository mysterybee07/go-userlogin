package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/userlogin/controllers"
)

// import "github.com/gofiber/fiber/v2"

// // import "github.com/gofiber/fiber/v2"

// // func Setup(app *fiber.App) error {
// // 	app.Get("/", controllers.Register)

// // }
// func Setup(app *fiber.App) {
// 	app.Get("/", controllers.Register)
// }

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
