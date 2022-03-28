package utils

import "github.com/gofiber/fiber/v2"

//GetJson
func GetJson(c *fiber.Ctx) {
	c.Set("Content-Type", fiber.MIMEApplicationJSON)
}
