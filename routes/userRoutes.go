package routes

import (
	"FIVERR/services"

	"github.com/gofiber/fiber"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/user_api")

	api.Get("/AllUsers", services.GetUsers)
	api.Get("/CustomUser/:id", services.GetUser)
	api.Post("/AddUser", services.NewUser)
	api.Delete("/DeleteUser/:id", services.DeleteUser)
	api.Put("/UpdateUser/:id", services.UpdateUser)
	api.Post("/RegisterUser", services.Register)
	api.Post("/Login", services.Login)
}
