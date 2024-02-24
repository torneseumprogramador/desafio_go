package routes

import (
	"http_fiber_api/app/controllers"
	"http_fiber_api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(app *fiber.App) {
	loginController := controllers.LoginController{}
	app.Post("/login", loginController.Login)

	homeController := controllers.HomeController{}
	app.Get("/", homeController.Index)

	recursosController := controllers.RecursosController{}
	app.Get("/recursos", middleware.ProtectedJwtLowLevel(), recursosController.Index)
	app.Post("/recursos", middleware.Protected(), recursosController.Create)
	app.Get("/recursos/:id", middleware.Protected(), recursosController.Show)
	app.Put("/recursos/:id", middleware.Protected(), recursosController.Update)
	app.Delete("/recursos/:id", middleware.Protected(), recursosController.Delete)
}
