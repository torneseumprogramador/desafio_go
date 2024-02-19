package main

import (
	"http_gin/src/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Carrega os templates HTML
	r.LoadHTMLGlob("src/templates/**/*.tmpl.html")

	r.Static("/public/", "./public_assets")

	config.Routes(r)

	r.Run(":5000") // Por padrão, escuta na porta 5000
}
