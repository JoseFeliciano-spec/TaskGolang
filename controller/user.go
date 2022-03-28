package controller

import (
	"TaskGolang/database"
	"TaskGolang/models"
	"TaskGolang/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(c *fiber.Ctx) error {
	utils.GetJson(c)

	var data models.UserLogin
	var dataQuery models.User

	c.BodyParser(&data)

	db, _ := database.Connection()

	db.Model(&models.User{}).Find(&dataQuery, "email = ?", data.Email)

	if data.Email != dataQuery.Email {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "No hay nada",
		})
	}

	errors := bcrypt.CompareHashAndPassword(dataQuery.Password, data.Password)
	if errors != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Falso")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    fmt.Sprint(dataQuery.ID),
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON("Verdad")
}

func UserRegister(c *fiber.Ctx) error {
	utils.GetJson(c)

	var data models.User

	c.BodyParser(&data)
	db, _ := database.Connection()

	fmt.Print(data)

	password, _ := bcrypt.GenerateFromPassword(data.Password, bcrypt.DefaultCost)

	msg := db.Create(&models.User{
		Name:     data.Name,
		LastName: data.LastName,
		Password: password,
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
