package routes

import (
	"go-api/controllers"
	"go-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {

	app.Post("/login", middlewares.AuthLogin, controllers.Login)

}
