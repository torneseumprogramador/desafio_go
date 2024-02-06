package main

import "fmt"

type PessoaFisica struct {
	Nome string
	CPF  string
}

func (pf PessoaFisica) ImprimeNome() {
	fmt.Println("Nome:", pf.Nome)
}

type PessoaFisicaNova struct {
	PessoaFisica // Composição, ao invés de herança
	Endereco     string
}

func main() {
	// Em Go, a "herança" é realizada através de um conceito chamado de composição.

	p := PessoaFisicaNova{
		PessoaFisica: PessoaFisica{
			Nome: "Teste",
			CPF:  "3232233223",
		},
		Endereco: "Endereço aqui",
	}

	fmt.Println(p.Nome) // Acesso direto graças à composição
	fmt.Println(p.CPF)
	fmt.Println(p.Endereco)

	p.ImprimeNome()
}
