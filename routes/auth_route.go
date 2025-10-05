package routes

import (
	"go-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {

	app.Post("/login", controllers.Login)

}
