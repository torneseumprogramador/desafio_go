package main

import (
	"fmt"

	"github.com/torneseumprogramador/validadores_golang"
)

func main() {

	fmt.Println("===============[CPF]================")

	cpf := "748.469.830-05"
	esperado := true
	resultado, err := validadores_golang.ValidarCPF(cpf)
	if resultado != esperado {
		fmt.Printf("ValidarCPF(%s) = %v; esperado %v (erro: %v)\n", cpf, resultado, esperado, err)
	} else {
		fmt.Printf("CPF (%s) validado com sucesso\n", cpf)
	}

	cpf = "748.469.830-01"
	resultado, err = validadores_golang.ValidarCPF(cpf)
	fmt.Printf("ValidarCPF(%s) = %v; (erro: %v)\n", cpf, resultado, err)

	fmt.Println("===============[CNPJ]================")

	cnpj := "25.823.284/0001-47"
	esperadoCnpj := true
	resultadoCnpj, errCnpj := validadores_golang.ValidarCNPJ(cnpj)
	if resultadoCnpj != esperadoCnpj {
		fmt.Printf("ValidarCNPJ(%s) = %v; esperado %v (erro: %v)\n", cnpj, resultadoCnpj, esperadoCnpj, errCnpj)
	} else {
		fmt.Printf("CNPJ (%s) validado com sucesso\n", cnpj)
	}

	cnpj = "25.823.284/0001-48"
	resultadoCnpj, errCnpj = validadores_golang.ValidarCNPJ(cnpj)
	fmt.Printf("ValidarCNPJ(%s) = %v; (erro: %v)\n", cpf, resultadoCnpj, errCnpj)
}
