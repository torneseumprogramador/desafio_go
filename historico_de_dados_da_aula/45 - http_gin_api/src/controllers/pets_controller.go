package controllers

import (
	"fmt"
	"http_gin/src/DTOs"
	"http_gin/src/database"
	"http_gin/src/models"
	"http_gin/src/repositorios"
	"http_gin/src/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func petRepo() *repositorios.PetRepositorioMySql {
	db, _ := database.GetDB()
	return &repositorios.PetRepositorioMySql{DB: db}
}

type PetsController struct{}

func (pc *PetsController) Index(c *gin.Context) {
	servico := servicos.NovoPetServico(petRepo())
	pets, erro := servico.ListaPetView()

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
	}

	c.JSON(http.StatusOK, pets)
}

func (pc *PetsController) Cadastrar(c *gin.Context) {
	var petDTO DTOs.PetDTO

	if err := c.BindJSON(&petDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet := models.Pet{
		Id:     "",
		Nome:   petDTO.Nome,
		DonoId: petDTO.DonoId,
		Tipo:   petDTO.Tipo,
	}

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	id, erro := servico.Repo.Adicionar(pet)
	pet.Id = id

	if erro == nil {
		c.JSON(http.StatusCreated, pet)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"erro": erro.Error(),
	})
}

func (pc *PetsController) Mostrar(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	petDb, erro := servico.Repo.BuscarPorId(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	if petDb == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": fmt.Errorf("pet não encontrado"),
		})
		return
	}

	c.JSON(http.StatusOK, petDb)
}

func (pc *PetsController) Excluir(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	servico.Repo.Excluir(id)

	c.JSON(http.StatusNoContent, gin.H{})
}

func (pc *PetsController) Alterar(c *gin.Context) {
	id := c.Param("id")
	var petDTO DTOs.PetDTO

	if err := c.BindJSON(&petDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	petDb, erro := servico.Repo.BuscarPorId(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	if petDb == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": fmt.Errorf("pet não encontrado"),
		})
		return
	}

	petDb.Nome = petDTO.Nome
	petDb.DonoId = petDTO.DonoId
	petDb.Tipo = petDTO.Tipo

	erroAlterar := servico.Repo.Alterar(*petDb)

	if erroAlterar == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erroAlterar.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, petDb)
}
