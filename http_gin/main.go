package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Carrega os templates HTML
	// r.LoadHTMLGlob("templates/*.tmpl.html")
	// r.LoadHTMLGlob("templates/**/*")
	r.LoadHTMLGlob("templates/**/*.tmpl.html")

	r.Static("/public/", "./public_assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl.html", gin.H{
			"title": "Página Principal",
		})
	})

	r.Run(":5000") // Por padrão, escuta na porta 8080
}
