package controllers

import (
	"html/template"
	"http_gin/src/enums"
	"http_gin/src/libs"
	"http_gin/src/models"
	"http_gin/src/servicos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PetsIndex(c *gin.Context) {
	pets, _ := servicos.ListaDePetsJson()

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Pets",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/index.tmpl.html",
					map[string][]models.Pet{
						"pets": pets,
					},
				),
			),
		},
	)
}

func PetsNovo(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Registro de Pet",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/salvar.tmpl.html",
					map[string]interface{}{
						"pet":    models.Pet{},
						"titulo": "Registro de Pet",
						"action": "/pets/cadastrar",
					},
				),
			),
		},
	)
}

func PetsCadastrar(c *gin.Context) {
	tipoInt, _ := strconv.Atoi(c.Request.FormValue("tipo"))

	pet := models.Pet{
		Id:   "",
		Nome: c.Request.FormValue("nome"),
		Dono: c.Request.FormValue("dono"),
		Tipo: enums.Tipo(tipoInt),
	}

	erro := servicos.AdicionarJson(pet)

	if erro == nil {
		c.Redirect(302, "/pets")
		return
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Registro de Pet",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   erro,
						"pet":    pet,
						"titulo": "Registro de Pet",
						"action": "/pets/cadastrar",
					},
				),
			),
		},
	)
}

func PetsExcluir(c *gin.Context) {
	id := c.Param("id")

	servicos.ExcluirJson(id)
	c.Redirect(302, "/pets")
}

func PetsEditar(c *gin.Context) {
	pet := servicos.BuscarPorIdJson(c.Param("id"))

	if pet == nil {
		c.Redirect(302, "/pets")
		return
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Alterando Pet",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   nil,
						"pet":    pet,
						"titulo": "Alterando um Pet",
						"action": "/pets/" + pet.Id + "/alterar",
					},
				),
			),
		},
	)
}

func PetsAlterar(c *gin.Context) {
	pet := servicos.BuscarPorIdJson(c.Param("id"))

	if pet == nil {
		c.Redirect(302, "/pets")
		return
	}

	pet.Nome = c.Request.FormValue("nome")
	pet.Dono = c.Request.FormValue("dono")

	tipoInt, _ := strconv.Atoi(c.Request.FormValue("tipo"))
	pet.Tipo = enums.Tipo(tipoInt)

	erro := servicos.AlterarJson(*pet)

	if erro == nil {
		c.Redirect(302, "/pets")
		return
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Alterando um Pet",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   erro,
						"pet":    pet,
						"titulo": "Alterando um Pet",
						"action": "/pets/" + pet.Id + "/alterar",
					},
				),
			),
		},
	)
}
