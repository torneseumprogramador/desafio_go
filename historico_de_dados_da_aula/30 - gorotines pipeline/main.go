package main

import (
	"fmt"
	"time"
)

// Estágio 1: Geração de Números
func generateNumbers(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			fmt.Println("generateNumbers", i)
			time.Sleep(time.Millisecond * 500)
			out <- i
		}
		close(out)
	}()
	return out
}

// Estágio 2: Processamento (multiplica cada número por 2)
func processNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Println("processNumbers n", n, " * 2 = ", n*2)
			time.Sleep(time.Millisecond * 500)
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// Estágio 3: Coleta e Exibição
func collectAndDisplay(in <-chan int) {
	for n := range in {
		fmt.Printf("Processado: %d\n", n)
	}
}

func main() {
	// Configura o pipeline
	numbers := generateNumbers(5)        // Estágio 1: Gera números de 1 a 5
	processed := processNumbers(numbers) // Estágio 2: Processa os números

	// Inicia o estágio final que coleta e exibe os resultados
	collectAndDisplay(processed) // Estágio 3

	fmt.Println("Pipeline concluído")
}
