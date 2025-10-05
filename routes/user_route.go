package routes

import (
	"go-api/controllers"
	"go-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {

	app.Post("/user", middlewares.UserValidation, controllers.Create)

	app.Get("/users", controllers.Get)

	app.Get("/user/:id", controllers.GetById)

}

func ProductRoute(app *fiber.App) {

	app.Post("/product", controllers.CreateProduct)

}
