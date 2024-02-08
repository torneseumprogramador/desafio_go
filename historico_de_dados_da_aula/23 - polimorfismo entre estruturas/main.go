////// Não existe polimorfismo entre estruturas, vc só consegue fazer polimorfismo com interfaces, como no arquivo: "main.go uso de implementaçao de contrato com interface" ////////

////// este código comentado da erro ////
// package main

// import "fmt"

// type PessoaFisica struct {
// 	Nome string
// 	CPF  string
// }

// func (pf PessoaFisica) ImprimeNome() {
// 	fmt.Println("Nome:", pf.Nome)
// }

// type PessoaFisicaNova struct {
// 	PessoaFisica // Composição, ao invés de herança
// 	Endereco     string
// }

// func (pf PessoaFisica) ImprimeNome() { // polimorfismo
// 	fmt.Println("%s, %s", pf.Nome, pf.Endereco)
// }

// func main() {
// 	// Em Go, a "herança" é realizada através de um conceito chamado de composição.

// 	p := PessoaFisicaNova{
// 		PessoaFisica: PessoaFisica{
// 			Nome: "Teste",
// 			CPF:  "3232233223",
// 		},
// 		Endereco: "Endereço aqui",
// 	}

// 	fmt.Println(p.Nome) // Acesso direto graças à composição
// 	fmt.Println(p.CPF)
// 	fmt.Println(p.Endereco)

// 	p.ImprimeNome()
// }

// / a forma é criar métodos com outros nomes, como abaixo
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

// Método específico para PessoaFisicaNova
func (pfn PessoaFisicaNova) ImprimeNomeComEndereco() {
	fmt.Printf("Nome: %s, Endereço: %s\n", pfn.Nome, pfn.Endereco)
}

func main() {
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

	p.ImprimeNome()            // Chama o método de PessoaFisica
	p.ImprimeNomeComEndereco() // Chama o método específico de PessoaFisicaNova
}
