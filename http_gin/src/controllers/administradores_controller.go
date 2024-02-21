package controllers

import (
	"fmt"
	"http_gin/src/DTOs"
	"http_gin/src/database"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"http_gin/src/repositorios"
	"http_gin/src/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func admRepositorio() *repositorios.AdministradorRepositorioMySql {
	db, _ := database.GetDB()
	return &repositorios.AdministradorRepositorioMySql{DB: db}
}

type AdministradoresController struct{}

func (pc *AdministradoresController) Index(c *gin.Context) {
	repo := admRepositorio()
	administradores, erro := repo.ListaAdmView()

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
	}

	c.JSON(http.StatusOK, administradores)
}

func (pc *AdministradoresController) Mostrar(c *gin.Context) {
	id := c.Param("id")

	repo := admRepositorio()
	admDb, erro := repo.BuscarPorIdModelView(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	if admDb == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": fmt.Errorf("pet não encontrado"),
		})
		return
	}

	c.JSON(http.StatusOK, admDb)
}

func (pc *AdministradoresController) Cadastrar(c *gin.Context) {
	var admDTO DTOs.AdministradorDTO

	if err := c.BindJSON(&admDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adm := models.Administrador{
		Id:    "",
		Nome:  admDTO.Nome,
		Email: admDTO.Email,
		Senha: admDTO.Senha,
		Super: admDTO.Super,
	}

	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
	id, erro := servico.Repo.Adicionar(adm)

	if erro == nil {
		c.JSON(http.StatusCreated, model_views.AdmView{
			Id:    id,
			Nome:  adm.Nome,
			Email: adm.Email,
			Super: adm.Super,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"erro": erro.Error(),
	})
}

func (pc *AdministradoresController) Excluir(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
	servico.Repo.Excluir(id)

	c.JSON(http.StatusNoContent, gin.H{})
}

func (pc *AdministradoresController) Alterar(c *gin.Context) {
	id := c.Param("id")
	var admDTO DTOs.AdministradorDTO

	if err := c.BindJSON(&admDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
	admDb, erro := servico.Repo.BuscarPorId(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	if admDb == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": fmt.Errorf("pet não encontrado"),
		})
		return
	}

	admDb.Nome = admDTO.Nome
	admDb.Email = admDTO.Email
	admDb.Senha = admDTO.Senha
	admDb.Super = admDTO.Super

	erroAlterar := servico.Repo.Alterar(*admDb)

	if erroAlterar == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erroAlterar.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model_views.AdmView{
		Id:    admDb.Id,
		Nome:  admDb.Nome,
		Email: admDb.Email,
		Super: admDb.Super,
	})
}
