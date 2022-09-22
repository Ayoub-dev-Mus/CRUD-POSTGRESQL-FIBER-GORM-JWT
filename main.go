package main

import (
	"FIVERR/database"
	"FIVERR/routes"

	"github.com/gofiber/fiber"
)

func main() {
	database.InitDatabase()

	app := fiber.New()
	routes.SetupCategoryRoutes(app)
	routes.SetupUserRoutes(app)
	app.Listen(8800)

}
