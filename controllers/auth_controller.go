package controllers

import (
	"go-api/config"

	"go-api/services"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Invalid request payload",
			"error":       err.Error(),
		})
	}

	user, tokenString, err := services.SignWithEmailAndPassword(input.Name, input.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Preload("Products").First(&user, user.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to load user data",
			"error":   err.Error(),
		})
	}

	user.Password = ""


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   tokenString,
		"data": fiber.Map{
			"name":     user.Name,
			"email":    user.Email,
			"products": user.Products,
		},
	})
}
