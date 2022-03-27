package controller

import (
	"TaskGolang/database"
	"TaskGolang/models"
	"TaskGolang/utils"

	"github.com/gofiber/fiber/v2"
)

func UserLogin(c *fiber.Ctx) error {
	utils.GetJson(c)
	var data models.User

	c.BodyParser(&data)
	db, _ := database.Connection()

	msg := db.Create(&models.User{
		Name:     data.Name,
		LastName: data.LastName,
		Password: data.Password,
		Email:    data.Email,
	})

	if msg.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(
			models.Message{
				Message: "Se ha creado el usuario",
				Data:    "",
				Code:    fiber.StatusOK,
			},
		)
	}

	return c.Status(fiber.StatusConflict).JSON(
		models.Message{
			Message: "No se ha creado el usuario",
			Data:    "",
			Code:    fiber.StatusConflict,
		},
	)
}
