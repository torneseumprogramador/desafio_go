package main

import (
	"fmt"
	"time"
)

// Função para imprimir números
func printNumbers(name string, count int, done chan bool) {
	for i := 1; i <= count; i++ {
		fmt.Println(name, ":", i)
		// Espera um pouco para ver a alternância entre as goroutines
		time.Sleep(time.Millisecond * 500)
	}
	// Envia um sinal para o canal indicando que esta goroutine terminou
	done <- true // terminei a tarefa (escrevendo)
}

func main() {
	// Cria um canal de booleanos
	canalSheila := make(chan bool)
	canalLana := make(chan bool)

	// Inicia as goroutines, passando o canal como argumento
	go printNumbers("Sheila", 5, canalSheila)
	go printNumbers("Lana", 8, canalLana)

	// Espera receber dois sinais de conclusão das goroutines
	<-canalSheila // Await (lendo)
	<-canalLana   // Await (lendo)

	fmt.Println("Programa terminou")
}
