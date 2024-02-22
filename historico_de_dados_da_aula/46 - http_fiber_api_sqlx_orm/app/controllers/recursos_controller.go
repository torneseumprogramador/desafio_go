package controllers

import (
	"fmt"
	"http_fiber_api/app/models"

	"github.com/gofiber/fiber/v2"
)

type RecursosController struct{}

func (pc *RecursosController) Index(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"mensagem": "Informação obtida com sucesso",
	})
}

func (pc *RecursosController) Create(c *fiber.Ctx) error {
	recurso := new(models.Recurso)

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
}

func (pc *RecursosController) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	fmt.Println(id)

	recurso := new(models.Recurso)

	return c.Status(200).JSON(recurso)
}

func (pc *RecursosController) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	recurso := new(models.Recurso)

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
}

func (pc *RecursosController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	return c.Status(204).SendString("")
}
