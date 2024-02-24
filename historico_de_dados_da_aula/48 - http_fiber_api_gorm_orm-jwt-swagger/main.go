package main

import (
	"http_fiber_api/pkg/routes"
	"http_fiber_api/plataform/database"
	"http_fiber_api/plataform/migrations"

	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "http_fiber_api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Desafio de Golang
// @description Este é um app construido junto com os alunos do torne-se um programador, com objetivo em aprender a linguagem Golang e seus componentes
// @version 1.0
// @BasePath /
// @contact.name Danilo Aparecido
// @contact.url https://www.torneseumprogramador.com.br/cursos/desafio_go_lang
// @contact.email suporte@torneseumprogramador.com.br
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	migrations.Migrate(database.GetConnection())
	migrations.Seed()

	app := fiber.New()

	// Configure o middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Cuidado com o uso do '*' em ambiente de produção
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
	}))

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	routes.RegisterPublicRoutes(app)
	app.Listen(":3000")
}
