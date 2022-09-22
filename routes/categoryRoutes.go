package routes

import(
	"FIVERR/services"
	"github.com/gofiber/fiber")

func SetupRoutes(app *fiber.App) {
	api := app.Group("/category_api")

	api.Get("/AllCategory", services.GetCategorys)
	api.Get("/CustomCategory/:id", services.GetCategory)
	api.Post("/AddCategory", services.NewCategory)
	api.Delete("/DeleteCategory/:id", services.DeleteCategory)
	api.Put("/UpdateCategory/:id", services.UpdateCategory)
}
