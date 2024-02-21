package main

import (
	"http_gin/src/config"

	"github.com/gin-gonic/gin"
)

func startWebApp() {
	r := gin.Default()

	config.Routes(r)

	r.Run(":5000") // Por padrão, escuta na porta 5000
}

func main() {
	startWebApp()
}
