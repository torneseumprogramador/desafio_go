package config

import (
	"http_gin/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", controllers.HomeIndex)

	r.GET("/pets", controllers.PetsIndex)
	r.GET("/pets/novo", controllers.PetsNovo)
	r.POST("/pets/cadastrar", controllers.PetsCadastrar)
	r.GET("/pets/:id/excluir", controllers.PetsExcluir)
	r.GET("/pets/:id/editar", controllers.PetsEditar)
	r.POST("/pets/:id/alterar", controllers.PetsAlterar)
}
