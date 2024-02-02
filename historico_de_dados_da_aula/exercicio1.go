package exercicos

import "fmt"

var sobrenome = "teste"

func main() {
	var year int = 1998   // Colocando variável ano em 1998
	var age = 2024 - year // Calculo para descobrir idade
	var nome = "Danilo"   // Definindo nome para concatenar

	var message string = "Hello, Go!" // Mensagem padrão

	message += " I'm [" + nome + "]." // Contatenando variável string

	telefone := "(11) 3434-4333"
	// var telefone = "(11) 3434-4333"

	// Mostrando resultados na tela
	fmt.Println("Age:", age)
	fmt.Println("Message:", message)
	fmt.Println("Telefone:", telefone)
	fmt.Println("sobrenome:", sobrenome)
}
