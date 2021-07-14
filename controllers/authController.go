package controllers

import (
	"../database"
	"../models"
	"github.com/gofiber/fiber"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	// BodyParser binds the request body to a struct.
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// match password and password_confirm
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}
