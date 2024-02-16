package controllers

import (
	"fmt"
	"html/template"
	"http_gin/src/database"
	"http_gin/src/libs"
	"http_gin/src/models"
	"http_gin/src/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DonosController struct{}

func (pc *DonosController) Index(c *gin.Context) {
	db, _ := database.GetDB()
	ps := servicos.DonoServico{DB: db}
	donos, _ := ps.Lista()

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Donos",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/index.tmpl.html",
					map[string][]models.Dono{
						"donos": donos,
					},
				),
			),
		},
	)
}

func (pc *DonosController) Novo(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Registro de Dono",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/salvar.tmpl.html",
					map[string]interface{}{
						"dono":   models.Dono{},
						"titulo": "Registro de Dono",
						"action": "/donos/cadastrar",
					},
				),
			),
		},
	)
}

func (pc *DonosController) Cadastrar(c *gin.Context) {
	dono := models.Dono{
		Id:       "",
		Nome:     c.Request.FormValue("nome"),
		Telefone: c.Request.FormValue("telefone"),
	}

	db, _ := database.GetDB()
	ps := servicos.DonoServico{DB: db}
	erro := ps.Adicionar(dono)

	if erro == nil {
		c.Redirect(302, "/donos")
		return
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Registro de Dono",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   erro,
						"dono":   dono,
						"titulo": "Registro de Dono",
						"action": "/donos/cadastrar",
					},
				),
			),
		},
	)
}

func (pc *DonosController) Excluir(c *gin.Context) {
	id := c.Param("id")

	db, _ := database.GetDB()
	ps := servicos.DonoServico{DB: db}
	ps.Excluir(id)
	c.Redirect(302, "/donos")
}

func (pc *DonosController) Editar(c *gin.Context) {
	db, _ := database.GetDB()
	ps := servicos.DonoServico{DB: db}
	dono, erro := ps.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/donos")
		return
	}

	if dono == nil {
		c.Redirect(302, "/donos")
		return
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Alterando Dono",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   nil,
						"dono":   dono,
						"titulo": "Alterando um Dono",
						"action": "/donos/" + dono.Id + "/alterar",
					},
				),
			),
		},
	)
}

func (pc *DonosController) Alterar(c *gin.Context) {
	db, _ := database.GetDB()
	ps := servicos.DonoServico{DB: db}
	dono, erro := ps.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/donos")
		return
	}

	if dono == nil {
		c.Redirect(302, "/donos")
		return
	}

	dono.Nome = c.Request.FormValue("nome")
	dono.Telefone = c.Request.FormValue("telefone")

	erroAlterar := ps.Alterar(*dono)

	if erroAlterar == nil {
		fmt.Println("Erro ao executar instrução sql ", erroAlterar.Error())
		c.Redirect(302, "/donos")
		return
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Alterando um Dono",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   erro,
						"dono":   dono,
						"titulo": "Alterando um Dono",
						"action": "/donos/" + dono.Id + "/alterar",
					},
				),
			),
		},
	)
}
