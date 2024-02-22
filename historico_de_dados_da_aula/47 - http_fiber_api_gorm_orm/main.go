package main

import (
	"fmt"
	"http_fiber_api/pkg/routes"
	"http_fiber_api/plataform/database"
	"http_fiber_api/plataform/migrations"

	"github.com/gofiber/fiber/v2"
)

func main() {
	migrations.Migrate(database.GetConnection())
	fmt.Println("==== Db Migrate com sucesso ====")

	app := fiber.New()
	routes.RegisterPublicRoutes(app)
	app.Listen(":3000")
}
