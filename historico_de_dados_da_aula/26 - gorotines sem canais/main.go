package main

import (
	"fmt"
	"time"
)

// Função para imprimir números
func printNumbers(name string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println(name, ":", i)
		// Espera um pouco para ver a alternância entre as goroutines
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// Inicia a primeira goroutine
	go printNumbers("Lana", 5)

	// Inicia a segunda goroutine
	go printNumbers("Danilo", 5)

	// Inicia a segunda goroutine
	go printNumbers("Sheila", 5)

	// Espera tempo suficiente para que as goroutines terminem
	time.Sleep(time.Second * 6) // -> Canais

	fmt.Println("Programa terminou")
}
