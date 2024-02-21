package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (hc *HomeController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Bem-vindo à API com Golang e Gin",
	})
}
