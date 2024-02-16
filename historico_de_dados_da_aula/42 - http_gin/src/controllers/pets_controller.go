package controllers

import (
	"html/template"
	"http_gin/src/enums"
	"http_gin/src/libs"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"http_gin/src/servicos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PetsController struct{}

func (pc *PetsController) Index(c *gin.Context) {
	ps := servicos.PetServico{}
	pets, _ := ps.ListaPetView(servicos.DonoServico{})

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Pets",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/index.tmpl.html",
					map[string][]model_views.PetView{
						"pets": pets,
					},
				),
			),
		},
	)
}

func donos() []models.Dono {
	ds := servicos.DonoServico{}
	donos, _ := ds.Lista()
	return donos
}

func (pc *PetsController) Novo(c *gin.Context) {
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
						"donos":  donos(),
					},
				),
			),
		},
	)
}

func (pc *PetsController) Cadastrar(c *gin.Context) {
	tipoInt, _ := strconv.Atoi(c.Request.FormValue("tipo"))

	pet := models.Pet{
		Id:     "",
		Nome:   c.Request.FormValue("nome"),
		DonoId: c.Request.FormValue("dono_id"),
		Tipo:   enums.Tipo(tipoInt),
	}

	ps := servicos.PetServico{}
	erro := ps.Adicionar(pet)

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
						"donos":  donos(),
					},
				),
			),
		},
	)
}

func (pc *PetsController) Excluir(c *gin.Context) {
	id := c.Param("id")

	ps := servicos.PetServico{}
	ps.Excluir(id)
	c.Redirect(302, "/pets")
}

func (pc *PetsController) Editar(c *gin.Context) {
	ps := servicos.PetServico{}
	pet := ps.BuscarPorId(c.Param("id"))

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
						"donos":  donos(),
					},
				),
			),
		},
	)
}

func (pc *PetsController) Alterar(c *gin.Context) {
	ps := servicos.PetServico{}
	pet := ps.BuscarPorId(c.Param("id"))

	if pet == nil {
		c.Redirect(302, "/pets")
		return
	}

	pet.Nome = c.Request.FormValue("nome")
	pet.DonoId = c.Request.FormValue("dono_id")

	tipoInt, _ := strconv.Atoi(c.Request.FormValue("tipo"))
	pet.Tipo = enums.Tipo(tipoInt)

	erro := ps.Alterar(*pet)

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
						"donos":  donos(),
					},
				),
			),
		},
	)
}
