package main

import (
	"fmt"
	"go-api/config"
	"go-api/migrate"
	"go-api/routes"

	"os"

	"github.com/gofiber/fiber/v2"

)

func main() {

	config.LoadEnv()

	migrate.Migrate()

	app := fiber.New()

	routes.UserRoute(app)
	routes.ProductRoute(app)

	routes.AuthRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Server running on port " + port)
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
