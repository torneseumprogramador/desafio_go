package main

import (
	"fmt"
	"time"
)

// Função para imprimir números e retornar um slice com os números
func printNumbers(name string, count int, done chan []int) {
	var nums []int // Slice para armazenar os números
	for i := 1; i <= count; i++ {
		fmt.Println(name, ": gravando no array o indice -> ", i)
		nums = append(nums, i) // Adiciona o número ao slice
		// Espera um pouco para ver a alternância entre as goroutines
		time.Sleep(time.Millisecond * 500)
	}
	// Envia o slice para o canal indicando que esta goroutine terminou
	done <- nums // Envia os dados processados
}

func main() {
	// Cria um canal de slices de inteiros
	canalSheila := make(chan []int)
	canalLana := make(chan []int)

	// Inicia as goroutines, passando o canal como argumento
	go printNumbers("Sheila", 5, canalSheila)
	go printNumbers("Lana", 8, canalLana)

	// Espera receber os dados processados das goroutines
	numsSheila := <-canalSheila // Recebe os dados processados
	numsLana := <-canalLana     // Recebe os dados processados

	fmt.Println("Números processados por Sheila:", numsSheila)
	fmt.Println("Números processados por Lana:", numsLana)

	fmt.Println("Programa terminou")
}
