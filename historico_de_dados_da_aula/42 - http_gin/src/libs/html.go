package libs

import (
	"bytes"
	"log"
	"text/template"
)

func Render(templateFile string, data interface{}) string {
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
