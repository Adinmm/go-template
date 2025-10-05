package migrate

import (
	"go-api/config"
	"go-api/models"
)

func Migrate() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
}