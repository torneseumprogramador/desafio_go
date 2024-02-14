package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func render(templateFile string, data interface{}) string {
	// Carrega o arquivo de template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatalf("Error parsing template files: %v", err)
	}

	// Cria um buffer para armazenar a saída do template
	var tmplBytes bytes.Buffer

	// Executa o template com os dados fornecidos e escreve a saída no buffer
	err = tmpl.Execute(&tmplBytes, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// Retorna a string resultante
	return tmplBytes.String()
}

func main() {
	r := gin.Default()

	// Carrega os templates HTML
	// r.LoadHTMLGlob("templates/*.tmpl.html")
	// r.LoadHTMLGlob("templates/**/*")
	r.LoadHTMLGlob("templates/**/*.tmpl.html")

	r.Static("/public/", "./public_assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl.html", gin.H{
			"title":        "Página Principal",
			"currentRoute": "home",
			"content":      template.HTML(render("templates/pages/home/index.tmpl.html", map[string]string{})),
		})
	})

	r.GET("/pets", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl.html", gin.H{
			"title":        "Pets",
			"currentRoute": "pets",
			"content": template.HTML(render("templates/pages/pets/index.tmpl.html", map[string]string{
				"Title": "Minha Página Inicial",
			})),
		})
	})

	r.Run(":5000") // Por padrão, escuta na porta 8080
}
