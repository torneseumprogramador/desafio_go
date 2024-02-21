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

func donoRepo() *repositorios.DonoRepositorioMySql {
	db, _ := database.GetDB()
	return &repositorios.DonoRepositorioMySql{DB: db}
}

type DonosController struct{}

func (pc *DonosController) Index(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	donos, erro := servico.Repo.Lista()

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
	}

	c.JSON(http.StatusOK, donos)
}

func (pc *DonosController) Mostrar(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	donoDb, erro := servico.Repo.BuscarPorId(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	if donoDb == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": fmt.Errorf("pet não encontrado"),
		})
		return
	}

	c.JSON(http.StatusOK, donoDb)
}

func (pc *DonosController) Cadastrar(c *gin.Context) {
	var donoDTO DTOs.DonoDTO

	if err := c.BindJSON(&donoDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dono := models.Dono{
		Id:       "",
		Nome:     donoDTO.Nome,
		Telefone: donoDTO.Telefone,
	}

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	id, erro := servico.Repo.Adicionar(dono)
	dono.Id = id

	if erro == nil {
		c.JSON(http.StatusCreated, dono)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"erro": erro.Error(),
	})
}

func (pc *DonosController) Excluir(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	servico.Repo.Excluir(id)

	c.JSON(http.StatusNoContent, gin.H{})
}

func (pc *DonosController) Alterar(c *gin.Context) {
	id := c.Param("id")
	var donoDTO DTOs.DonoDTO

	if err := c.BindJSON(&donoDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	donoDb, erro := servico.Repo.BuscarPorId(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	if donoDb == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": fmt.Errorf("pet não encontrado"),
		})
		return
	}

	donoDb.Nome = donoDTO.Nome
	donoDb.Telefone = donoDTO.Telefone

	erroAlterar := servico.Repo.Alterar(*donoDb)

	if erroAlterar == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erroAlterar.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, donoDb)
}
