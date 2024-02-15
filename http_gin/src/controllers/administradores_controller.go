package controllers

import (
	"html/template"
	"http_gin/src/libs"
	"http_gin/src/models"
	"http_gin/src/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdministradoresController struct{}

func (pc *AdministradoresController) Index(c *gin.Context) {
	ps := servicos.AdministradorServico{}
	administradores, _ := ps.Lista()

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Administradores",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/index.tmpl.html",
					map[string][]models.Administrador{
						"administradores": administradores,
					},
				),
			),
		},
	)
}

func (pc *AdministradoresController) Novo(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Registro de Administrador",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/salvar.tmpl.html",
					map[string]interface{}{
						"administrador": models.Administrador{},
						"titulo":        "Registro de Administrador",
						"action":        "/administradores/cadastrar",
					},
				),
			),
		},
	)
}

func (pc *AdministradoresController) Cadastrar(c *gin.Context) {
	administrador := models.Administrador{
		Id:    "",
		Nome:  c.Request.FormValue("nome"),
		Email: c.Request.FormValue("email"),
		Senha: c.Request.FormValue("senha"),
	}

	ps := servicos.AdministradorServico{}
	erro := ps.Adicionar(administrador)

	if erro == nil {
		c.Redirect(302, "/administradores")
		return
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Registro de Administrador",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/salvar.tmpl.html",
					map[string]interface{}{
						"erro":          erro,
						"administrador": administrador,
						"titulo":        "Registro de Administrador",
						"action":        "/administradores/cadastrar",
					},
				),
			),
		},
	)
}

func (pc *AdministradoresController) Excluir(c *gin.Context) {
	id := c.Param("id")

	ps := servicos.AdministradorServico{}
	ps.Excluir(id)
	c.Redirect(302, "/administradores")
}

func (pc *AdministradoresController) Editar(c *gin.Context) {
	ps := servicos.AdministradorServico{}
	administrador := ps.BuscarPorId(c.Param("id"))

	if administrador == nil {
		c.Redirect(302, "/administradores")
		return
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Alterando Administrador",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/salvar.tmpl.html",
					map[string]interface{}{
						"erro":          nil,
						"administrador": administrador,
						"titulo":        "Alterando um Administrador",
						"action":        "/administradores/" + administrador.Id + "/alterar",
					},
				),
			),
		},
	)
}

func (pc *AdministradoresController) Alterar(c *gin.Context) {
	ps := servicos.AdministradorServico{}
	administrador := ps.BuscarPorId(c.Param("id"))

	if administrador == nil {
		c.Redirect(302, "/administradores")
		return
	}

	administrador.Nome = c.Request.FormValue("nome")
	administrador.Email = c.Request.FormValue("email")
	administrador.Senha = c.Request.FormValue("senha")

	erro := ps.Alterar(*administrador)

	if erro == nil {
		c.Redirect(302, "/administradores")
		return
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Alterando um Administrador",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/salvar.tmpl.html",
					map[string]interface{}{
						"erro":          erro,
						"administrador": administrador,
						"titulo":        "Alterando um Administrador",
						"action":        "/administradores/" + administrador.Id + "/alterar",
					},
				),
			),
		},
	)
}
