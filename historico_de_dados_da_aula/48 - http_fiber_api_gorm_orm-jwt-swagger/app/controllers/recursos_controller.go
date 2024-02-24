package controllers

import (
	"http_fiber_api/app/dtos"
	"http_fiber_api/app/models"
	"http_fiber_api/app/services"
	"http_fiber_api/plataform/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RecursosController struct{}

// Index godoc
// @Summary Lista todos os recursos
// @Description Retorna uma lista de todos os recursos
// @Tags Recursos
// @Accept json
// @Produce json
// @Success 200 {array} models.Recurso
// @Router /recursos [get]
// @Security Bearer
func (pc *RecursosController) Index(c *fiber.Ctx) error {

	var recursos []models.Recurso
	err := services.GetAllRecursos(database.GetConnection(), &recursos)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": err.Error(),
		})
	}

	return c.Status(200).JSON(recursos)
}

// Create godoc
// @Summary Cria um novo recurso
// @Description Cria um novo recurso com as informações fornecidas
// @Tags Recursos
// @Accept json
// @Produce json
// @Param recurso body dtos.RecursoDTO true "Informações do Recurso"
// @Success 201 {object} models.Recurso
// @Router /recursos [post]
// @Security Bearer
func (pc *RecursosController) Create(c *fiber.Ctx) error {
	recursoDto := new(dtos.RecursoDTO)

	if err := c.BodyParser(recursoDto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"mensagem": "Erro ao analisar o corpo da requisição - " + err.Error(),
		})
	}

	recurso := models.Recurso{
		Nome:  recursoDto.Nome,
		Valor: recursoDto.Valor,
	}

	err := services.CreateRecurso(database.GetConnection(), &recurso)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": err.Error(),
		})
	}

	return c.Status(201).JSON(recurso)
}

// Show godoc
// @Summary Mostra um recurso
// @Description Retorna detalhes de um recurso específico pelo ID
// @Tags Recursos
// @Accept json
// @Produce json
// @Param id path int true "ID do Recurso"
// @Success 200 {object} models.Recurso
// @Router /recursos/{id} [get]
// @Security Bearer
func (pc *RecursosController) Show(c *fiber.Ctx) error {
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

// Update godoc
// @Summary Atualiza um recurso
// @Description Atualiza o recurso com o ID fornecido
// @Tags Recursos
// @Accept json
// @Produce json
// @Param id path int true "ID do Recurso"
// @Param recurso body dtos.RecursoDTO true "Informações atualizadas do Recurso"
// @Success 200 {object} models.Recurso
// @Router /recursos/{id} [put]
// @Security Bearer
func (pc *RecursosController) Update(c *fiber.Ctx) error {
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

	recursoDto := new(dtos.RecursoDTO)

	if err := c.BodyParser(recursoDto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"mensagem": "Erro ao analisar o corpo da requisição - " + err.Error(),
		})
	}

	recursoDb.Nome = recursoDto.Nome
	recursoDb.Valor = recursoDto.Valor

	err = services.UpdateRecurso(database.GetConnection(), &recursoDb)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mensagem": err.Error(),
		})
	}

	return c.Status(200).JSON(recursoDb)
}

// Delete godoc
// @Summary Exclui um recurso
// @Description Exclui o recurso com o ID fornecido
// @Tags Recursos
// @Accept json
// @Produce json
// @Param id path int true "ID do Recurso"
// @Success 204
// @Router /recursos/{id} [delete]
// @Security Bearer
func (pc *RecursosController) Delete(c *fiber.Ctx) error {
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
		return c.Status(404).JSON(fiber.Map{
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
