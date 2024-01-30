package exercicios

import (
	"fmt"
	"strconv"
	"time"
)

var sobrenome = "teste"

func main() {
	var nome string
	var ano int
	var mesNascimento int

	fmt.Print("Digite seu nome: \n")
	fmt.Scanln(&nome) // Definindo nome dinamicamente

	fmt.Printf("Digite o ano de seu nascimento %s\n", nome)
	fmt.Scanln(&ano) // Definindo nome dinamicamente

	fmt.Printf("Beleza %s, para ter uma precisão em sua idade, digite por favor o mês de seu nascimento\n", nome)
	fmt.Scanln(&mesNascimento) // Definindo nome dinamicamente

	var idade = time.Now().Year() - ano    // Calculo para descobrir idade
	var mesAtual = int(time.Now().Month()) // Capturando o mes atual

	if mesAtual < mesNascimento {
		idade -= 1
	}

	var message string = "Segue os seus dados:\n" // Mensagem padrão

	message += "Nome: " + nome + "\n"                 // Contatenando variável string
	message += "Idade: " + strconv.Itoa(idade) + "\n" // Contatenando variável string

	fmt.Println(message)
}
