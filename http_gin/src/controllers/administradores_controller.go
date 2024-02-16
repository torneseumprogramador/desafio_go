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

func admServicoMysql() *servicos.AdministradorServicoMySql {
	db, _ := database.GetDB()
	return &servicos.AdministradorServicoMySql{DB: db}
}

// TODO -> proxima aula
// func admServicoJson() *servicos.AdministradorServicoMySql {
// 	db, _ := database.GetDB()
// 	return &servicos.AdministradorServicoMySql{DB: db}
// }

type AdministradoresController struct{}

func (pc *AdministradoresController) Index(c *gin.Context) {
	repo := servicos.NovoRepositorio[models.Administrador](admServicoMysql())

	administradores, _ := repo.Servico.Lista()
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

	repo := servicos.NovoRepositorio[models.Administrador](admServicoMysql())
	erro := repo.Servico.Adicionar(administrador)

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

	repo := servicos.NovoRepositorio[models.Administrador](admServicoMysql())
	repo.Servico.Excluir(id)
	c.Redirect(302, "/administradores")
}

func (pc *AdministradoresController) Editar(c *gin.Context) {
	repo := servicos.NovoRepositorio[models.Administrador](admServicoMysql())
	administrador, erro := repo.Servico.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/administradores")
		return
	}

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
	repo := servicos.NovoRepositorio[models.Administrador](admServicoMysql())
	administrador, erro := repo.Servico.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/administradores")
		return
	}

	if administrador == nil {
		c.Redirect(302, "/administradores")
		return
	}

	administrador.Nome = c.Request.FormValue("nome")
	administrador.Email = c.Request.FormValue("email")
	administrador.Senha = c.Request.FormValue("senha")

	erroAlterar := repo.Servico.Alterar(*administrador)

	if erroAlterar == nil {
		fmt.Println("Erro ao executar instrução sql ", erroAlterar.Error())
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
