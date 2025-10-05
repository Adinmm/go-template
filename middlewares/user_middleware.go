package middlewares

import (
	"go-api/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserValidation(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Invalid request payload",
			"error":       err.Error(),
		})
	}

	log.Println(len(input.Name))


	if len(input.Name) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Name must be at least 3 characters long",
		})
	}


	return c.Next()
}
