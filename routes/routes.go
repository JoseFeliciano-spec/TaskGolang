package routes

import (
	"TaskGolang/controller"
	"TaskGolang/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Routes
func Routes(app *fiber.App) {
	db, _ := database.Connection()
	fmt.Println(db)

	app.Post("api/login", func(c *fiber.Ctx) error {
		return c.JSON("hola")
	})

	app.Post("api/register", controller.UserLogin)
}
