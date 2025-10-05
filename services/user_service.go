package services

import (
	"errors"
	"go-api/models"
	"go-api/repositories"
)

type UserResponse struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	Email    string           `json:"email"`
	Products []models.Product `json:"products"`
}

func GetUsers() ([]models.User, []UserResponse, error) {

	result, err := repositories.GetAllUsers()
	var userResponse []UserResponse

	if err != nil {
		return nil, nil, errors.New("failed to get users")

	}

	return result,  userResponse, nil

}
