package main

import (
	"aula_go/modulos/aniversario"
	"aula_go/modulos/com_retorno"
	"aula_go/modulos/void"
	"fmt"
)

// PascalCase <- para funções de módulos
// camelCase <- para função interna
// snake_case underscore_case underline_case <- variáveis

func main() {

	nomeFormatado := com_retorno.FormataString("Danilo")

	fmt.Printf("%s", nomeFormatado)

	void.ImprimirValor()

	fmt.Println("--------------")

	var nome string
	var dataAniversario string

	fmt.Println("Digite seu nome")
	fmt.Scanln(&nome)

	fmt.Println("Digite sua data de aniversário no formato: dd/mm/aaaa")
	fmt.Scanln(&dataAniversario)

	nome, idade, diasFaltamAniversario, horasFaltamAniversario := aniversario.QuantosDiasFaltamParaSeuAniversario(nome, dataAniversario)
	fmt.Printf("%s, você tem %d anos e faltam %d dias e %d horas para o seu aniversário\n", nome, idade, diasFaltamAniversario, horasFaltamAniversario)
}
