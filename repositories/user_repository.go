package repositories

import (
	"go-api/config"
	"go-api/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	result := config.DB.Preload("Products").Find(&users)

	return users, result.Error

}
