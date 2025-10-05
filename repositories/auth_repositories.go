package repositories

import (
	"go-api/config"
	"go-api/models"
)

func FindByUsername(username string) (models.User, error) {
	var user models.User
	result := config.DB.Where("name = ?", username).First(&user)
	return user, result.Error
}
