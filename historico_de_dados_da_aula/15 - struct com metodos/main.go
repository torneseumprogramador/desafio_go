package main

import "fmt"

type Produto struct {
	Codigo    int
	Nome      string
	Descricao string
	Preco     float32
}

func (p Produto) Apresentar() {
	fmt.Printf("Id: %d\nNome: %s\nDescrição: %s\nPreço: %.2f\n", p.Codigo, p.Nome, p.Descricao, p.Preco)
}

type Soma struct {
	X int
	Y int
}

func (soma Soma) Acao() int {
	return soma.X + soma.Y
}

func main() {
	produto := Produto{
		Codigo:    1,
		Nome:      "Celular Ix445",
		Descricao: "Um celular bem legal",
		Preco:     5.7,
	}

	produto.Apresentar()

	r := Soma{X: 1, Y: 2}.Acao()
	fmt.Printf("O resultado de ação de soma é: %d\n", r)
}
