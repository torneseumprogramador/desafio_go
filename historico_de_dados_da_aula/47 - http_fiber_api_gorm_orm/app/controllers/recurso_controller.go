package controllers

import (
	"http_fiber_api/app/models"
	"http_fiber_api/app/services"
	"http_fiber_api/plataform/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RecursoController struct{}

func (pc *RecursoController) Index(c *fiber.Ctx) error {

	var recursos []models.Recurso
	err := services.GetAllRecursos(database.GetConnection(), &recursos)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": err.Error(),
		})
	}

	return c.Status(200).JSON(recursos)
}

func (pc *RecursoController) Create(c *fiber.Ctx) error {
	recurso := new(models.Recurso)

	if err := c.BodyParser(recurso); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"mensagem": "Erro ao analisar o corpo da requisição - " + err.Error(),
		})
	}

	err := services.CreateRecurso(database.GetConnection(), recurso)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": err.Error(),
		})
	}

	return c.Status(201).JSON(recurso)
}

func (pc *RecursoController) Show(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, errIdStr := strconv.ParseUint(idStr, 10, 32)
	if errIdStr != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": errIdStr.Error(),
		})
	}

	var recurso models.Recurso
	err := services.GetRecursoByID(database.GetConnection(), &recurso, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": err.Error(),
		})
	}

	return c.Status(200).JSON(recurso)
}

func (pc *RecursoController) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, errIdStr := strconv.ParseUint(idStr, 10, 32)
	if errIdStr != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": errIdStr.Error(),
		})
	}

	var recursoDb models.Recurso
	err := services.GetRecursoByID(database.GetConnection(), &recursoDb, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": err.Error(),
		})
	}

	recurso := new(models.Recurso)

	if err := c.BodyParser(recurso); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"mensagem": "Erro ao analisar o corpo da requisição - " + err.Error(),
		})
	}

	recursoDb.Nome = recurso.Nome
	recursoDb.Valor = recurso.Valor

	err = services.UpdateRecurso(database.GetConnection(), &recursoDb)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": err.Error(),
		})
	}

	return c.Status(200).JSON(recursoDb)
}

func (pc *RecursoController) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, errIdStr := strconv.ParseUint(idStr, 10, 32)
	if errIdStr != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": errIdStr.Error(),
		})
	}

	var recursoDb models.Recurso
	err := services.GetRecursoByID(database.GetConnection(), &recursoDb, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": err.Error(),
		})
	}

	errDelete := services.DeleteRecurso(database.GetConnection(), &recursoDb)

	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": errDelete.Error(),
		})
	}
	return c.Status(204).SendString("")
}
