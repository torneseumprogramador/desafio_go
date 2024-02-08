package main

import "fmt"

func main() {
	// var meuSlice []int // slice = array dinamico
	// fmt.Printf("A quantidade é %d\n", len(meuSlice))

	// meuSlice = append(meuSlice, 10) // adicionando 1 item no meu slice (array dinamico)
	// fmt.Printf("A quantidade é %d\n", len(meuSlice))

	// meuSlice = append(meuSlice, 20, 30) // adicionando 2 itens no meu slice (array dinamico)
	// fmt.Printf("A quantidade é %d\n", len(meuSlice))

	// meuSlice = append(meuSlice, 20, 30, 4, 5, 6) // adicionando 5 itens no meu slice (array dinamico)
	// fmt.Printf("A quantidade é %d\n", len(meuSlice))

	sliceComDadoInicial := make([]int, 0, 5) // Cria um slice de inteiros com tamanho 0 e capacidade 5
	fmt.Printf("A quantidade é %d\n", len(sliceComDadoInicial))

	sliceComDadoInicial = append(sliceComDadoInicial, 20, 30, 4, 5, 6, 7, 22) // adicionando 5 itens no meu slice (array dinamico)
	fmt.Printf("A quantidade é %d\n", len(sliceComDadoInicial))

	for i, valor := range sliceComDadoInicial {
		fmt.Printf("Índice %d, Valor %d\n", i, valor)
	}
}
