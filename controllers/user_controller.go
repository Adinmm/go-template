package controllers

import (
	"go-api/config"
	"go-api/models"
	"go-api/services"
	"go-api/utils"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var user models.User


	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Invalid request body",
			"error":       err.Error(),
		})
	}

	
	if user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Password is required",
		})
	}

	
	passwordHashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"message":     "Failed to hash password",
			"error":       err.Error(),
		})
	}

	user.Password = passwordHashed


	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"message":     "Failed to create user",
			"error":       result.Error.Error(),
		})
	}

	
	user.Password = ""

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status_code": fiber.StatusCreated,
		"message":     "User created successfully",
	})
}

func Get(c *fiber.Ctx) error {

	users, userResponse, err := services.GetUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"message":     "Failed to retrieve users",
			"error":       err.Error(),
		})
	}

	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status_code": fiber.StatusNotFound,
			"message":     "No users found",
		})
	}

	for _, result := range users {
		userResponse = append(userResponse, services.UserResponse{
			ID:       result.ID,
			Name:     result.Name,
			Email:    result.Email,
			Products: result.Products,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_code": fiber.StatusOK,
		"message":     "Users retrieved successfully",
		"data":        userResponse,
	})
}

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "ID user is required",
		})
	}

	var user models.User

	result := config.DB.Find(&user, id)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status_code": fiber.StatusNotFound,
			"message":     "User not found",
			"error":       result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_code": fiber.StatusOK,
		"message":     "User retrieved successfully",
		"data":        user,
	})
}

func CreateProduct(app *fiber.Ctx) error {
	var product models.Product
	if err := app.BodyParser(&product); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Invalid request payload",
			"error":       err.Error(),
		})
	}

	result := config.DB.Create(&product)

	if result.Error != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"message":     "Failed to create product",
			"error":       result.Error.Error(),
		})
	}

	return app.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status_code": fiber.StatusCreated,
		"message":     "Product created successfully",
		"data":        product,
	})

}
