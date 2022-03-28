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

	app.Get("api/verifyme", func(c *fiber.Ctx) error {
		c.Cookies("jwt")
		fmt.Println(c.Cookies("jwt"))
		return c.SendString(fmt.Sprint(c.Cookies("data")))
	})

	app.Post("api/login", controller.UserLogin)

	app.Post("api/register", controller.UserRegister)
}
