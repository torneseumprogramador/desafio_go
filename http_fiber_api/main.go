package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Recurso struct {
	Nome  string  `json:"nome"`
	Valor float32 `json:"valor"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"mensagem": "API com fiber",
		})
	})

	app.Get("/recurso", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"mensagem": "Informação obtida com sucesso",
		})
	})

	app.Post("/recurso", func(c *fiber.Ctx) error {
		recurso := new(Recurso)

		if err := c.BodyParser(recurso); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"mensagem": "Erro ao analisar o corpo da requisição - " + err.Error(),
			})
		}

		fmt.Println(recurso.Nome)
		fmt.Println(recurso.Valor)

		return c.Status(201).JSON(fiber.Map{
			"mensagem": "Recurso criado com sucesso",
			"recurso":  recurso,
		})
	})

	app.Get("/recurso/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		fmt.Println(id)

		recurso := new(Recurso)

		return c.Status(200).JSON(recurso)
	})

	app.Put("/recurso/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		recurso := new(Recurso)

		if err := c.BodyParser(recurso); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"mensagem": "Erro ao analisar o corpo da requisição - " + err.Error(),
			})
		}

		fmt.Println(recurso.Nome)
		fmt.Println(recurso.Valor)

		return c.Status(200).JSON(fiber.Map{
			"mensagem": "Recurso atualizado com sucesso",
			"id":       id,
		})
	})

	app.Delete("/recurso/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		fmt.Println(id)
		return c.Status(204).SendString("")
	})

	app.Listen(":3000")
}
