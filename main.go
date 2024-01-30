package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var sobrenome = "teste"

func main() {
	var nome string
	var dataNascimento string
	var diaNascimento int
	var mesNascimento int
	var anoNascimento int

	fmt.Print("Digite seu nome: \n")
	fmt.Scanln(&nome) // Definindo nome dinamicamente

	fmt.Printf("Digite a data de seu nascimento no formato dia/mes/ano %s\n", nome)
	fmt.Scanln(&dataNascimento) // Definindo nome dinamicamente

	var partes = strings.Split(dataNascimento, "/")
	diaNascimento, _ = strconv.Atoi(partes[0])
	mesNascimento, _ = strconv.Atoi(partes[1])
	anoNascimento, _ = strconv.Atoi(partes[2])

	var idade = time.Now().Year() - anoNascimento // Calculo para descobrir idade
	var mesAtual = int(time.Now().Month())        // Capturando o mes atual
	var diaAtual = int(time.Now().Day())          // Capturando o dia atual

	if mesAtual < mesNascimento {
		idade -= 1
	} else if diaAtual < diaNascimento {
		idade -= 1
	}

	var message string = "Segue os seus dados:\n" // Mensagem padrão

	message += "Nome: " + nome + "\n"                 // Contatenando variável string
	message += "Idade: " + strconv.Itoa(idade) + "\n" // Contatenando variável string

	fmt.Println(message)
}
