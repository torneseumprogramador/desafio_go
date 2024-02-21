package config

import (
	"http_gin/src/controllers"
	"http_gin/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	homeController := controllers.HomeController{}
	r.GET("/", homeController.Index)

	loginController := controllers.LoginController{}
	r.POST("/login", loginController.Login)

	protectedRoutes := r.Group("/").Use(middlewares.AuthRequired())
	{
		petsController := controllers.PetsController{}
		protectedRoutes.GET("/pets", petsController.Index)
		protectedRoutes.POST("/pets", petsController.Cadastrar)
		protectedRoutes.GET("/pets/:id", petsController.Mostrar)
		protectedRoutes.DELETE("/pets/:id", petsController.Excluir)
		protectedRoutes.PUT("/pets/:id", petsController.Alterar)

		donosController := controllers.DonosController{}
		protectedRoutes.GET("/donos", donosController.Index)
		protectedRoutes.POST("/donos", donosController.Cadastrar)
		protectedRoutes.DELETE("/donos/:id", donosController.Excluir)
		protectedRoutes.PUT("/donos/:id", donosController.Alterar)
		protectedRoutes.GET("/donos/:id", donosController.Mostrar)

		administradoresController := controllers.AdministradoresController{}
		protectedRoutes.GET("/administradores", administradoresController.Index)
		protectedRoutes.POST("/administradores", administradoresController.Cadastrar)
		protectedRoutes.DELETE("/administradores/:id", administradoresController.Excluir)
		protectedRoutes.PUT("/administradores/:id", administradoresController.Alterar)
		protectedRoutes.GET("/administradores/:id", administradoresController.Mostrar)
	}
}
