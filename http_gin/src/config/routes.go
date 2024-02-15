package config

import (
	"http_gin/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	homeController := controllers.HomeController{}
	r.GET("/", homeController.Index)

	petsController := controllers.PetsController{}
	r.GET("/pets", petsController.Index)
	r.GET("/pets/novo", petsController.Novo)
	r.POST("/pets/cadastrar", petsController.Cadastrar)
	r.GET("/pets/:id/excluir", petsController.Excluir)
	r.GET("/pets/:id/editar", petsController.Editar)
	r.POST("/pets/:id/alterar", petsController.Alterar)

	donosController := controllers.DonosController{}
	r.GET("/donos", donosController.Index)
	r.GET("/donos/novo", donosController.Novo)
	r.POST("/donos/cadastrar", donosController.Cadastrar)
	r.GET("/donos/:id/excluir", donosController.Excluir)
	r.GET("/donos/:id/editar", donosController.Editar)
	r.POST("/donos/:id/alterar", donosController.Alterar)

	administradoresController := controllers.AdministradoresController{}
	r.GET("/administradores", administradoresController.Index)
	r.GET("/administradores/novo", administradoresController.Novo)
	r.POST("/administradores/cadastrar", administradoresController.Cadastrar)
	r.GET("/administradores/:id/excluir", administradoresController.Excluir)
	r.GET("/administradores/:id/editar", administradoresController.Editar)
	r.POST("/administradores/:id/alterar", administradoresController.Alterar)
}
