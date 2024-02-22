package routes

import (
	"http_fiber_api/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(app *fiber.App) {
	homeController := controllers.HomeController{}
	app.Get("/", homeController.Index)

	recursoController := controllers.RecursoController{}
	app.Get("/recurso", recursoController.Index)
	app.Post("/recurso", recursoController.Create)
	app.Get("/recurso/:id", recursoController.Show)
	app.Put("/recurso/:id", recursoController.Update)
	app.Delete("/recurso/:id", recursoController.Delete)
}
