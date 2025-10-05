package services

import (
	"errors"
	"time" 
	"go-api/config"
	"go-api/models"
	"go-api/repositories"
	"go-api/utils"

	"github.com/golang-jwt/jwt/v5"
)


func SignWithEmailAndPassword(username, password string) (models.User, string, error) {
	user, err := repositories.FindByUsername(username)
	if err != nil {
		// Tidak bocorkan apakah username atau password salah
		return models.User{}, "", errors.New("invalid username or password")
	}


	if !utils.CheckPasswordHash(password, user.Password) {
		return models.User{}, "", errors.New("invalid username or password")
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
		"email":    user.Email,
		"uid":      user.ID,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(config.JWTSecret)

	if err != nil {
		return models.User{}, "", errors.New("failed to sign token")
	}

	return user, tokenString, nil
}
