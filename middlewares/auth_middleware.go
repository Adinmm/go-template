package middlewares

import (
	"go-api/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func AuthLogin(c *fiber.Ctx) error {

	var input = dto.LoginInput{}

	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Invalid request payload",
			"error":       err.Error(),
		})
	}

	if err := input.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
	
		})
	}

	return c.Next()

}
