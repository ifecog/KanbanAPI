package controllers

import (
	"courier-api/config"
	"courier-api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterCourier(c *fiber.Ctx) error {
	type Request struct {
		FirstName	string	`json:"first_name"`	
		LastName	string	`json:"last_name"`	
		Email		string	`json:"email"`
		Phone		string	`json:"phone"`
		Password	string	`json:"password"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	courier := models.Courier{
		FirstName: 	body.FirstName,
		LastName: 	body.LastName,
		Email: 		body.Email,
		Phone: 		body.Phone,
		Password:	string(hashedPassword),
	}

	config.DB.Create(&courier)
	return c.JSON(courier)
}