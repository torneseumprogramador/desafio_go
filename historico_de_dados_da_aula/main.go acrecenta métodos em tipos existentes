package main

import (
	"fmt"
)

type PessoaFisica struct {
	Nome string
	CPF  string
}

type PessoaFisicaNova PessoaFisica // Criando um tipo novo em base de um tipo existente para acrescentar métodos
func (p PessoaFisicaNova) MostraDocumento() {
	fmt.Printf("%s", p.CPF)
}

func main() {
	p := new(PessoaFisicaNova)
	p.Nome = "Teste"
	p.CPF = "Teste"
	p.MostraDocumento()
}
