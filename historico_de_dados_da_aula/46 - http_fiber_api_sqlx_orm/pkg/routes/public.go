package routes

import (
	"http_fiber_api/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(app *fiber.App) {
	homeController := controllers.HomeController{}
	app.Get("/", homeController.Index)

	recursosController := controllers.RecursosController{}
	app.Get("/recurso", recursosController.Index)
	app.Post("/recurso", recursosController.Create)
	app.Get("/recurso/:id", recursosController.Show)
	app.Put("/recurso/:id", recursosController.Update)
	app.Delete("/recurso/:id", recursosController.Delete)
}
