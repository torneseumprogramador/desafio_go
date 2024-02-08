package main

import (
	"fmt"
	"time"
)

// Define uma estrutura para as informações a serem retornadas
type info struct {
	Name   string
	Nums   []int
	Status string
}

// Função para imprimir números e retornar informações processadas
func printNumbers(name string, count int, done chan info, contextoCompartilhado []int, next chan []int) {
	var nums []int
	for i := 1; i <= count; i++ {
		fmt.Println(name, ":", i)
		nums = append(nums, i)

		if contextoCompartilhado != nil {
			fmt.Println("---------------------------")
			fmt.Println("Dados compartilhados:", contextoCompartilhado)
		}
		time.Sleep(time.Millisecond * 500)
	}

	// Compartilha o contexto com outra goroutine, se necessário
	if next != nil {
		next <- nums // Passa os números processados para a próxima goroutine
	}

	// Envia as informações processadas para o canal done
	done <- info{Name: name, Nums: nums, Status: "completed"}
}

func main() {
	// Cria canais para as informações e contexto compartilhado
	canalInfoSheila := make(chan info)
	canalInfoLana := make(chan info)
	canalContexto := make(chan []int)

	// Inicia a goroutine de Lana, que compartilhará seu contexto com Sheila
	go printNumbers("Lana", 8, canalInfoLana, nil, canalContexto)

	// Espera pelo contexto de Lana antes de iniciar Sheila
	numsFromLana := <-canalContexto

	// Inicia a goroutine de Sheila, passando o contexto de Lana
	go printNumbers("Sheila", 5, canalInfoSheila, numsFromLana, nil) // Sheila não passa seu contexto adiante

	// Recebe as informações processadas de ambas as goroutines
	infoLana := <-canalInfoLana
	infoSheila := <-canalInfoSheila

	fmt.Printf("Informações de Lana: %+v\n", infoLana)
	fmt.Printf("Informações de Sheila: %+v, usando contexto de Lana: %v\n", infoSheila, numsFromLana)

	fmt.Println("Programa terminou")
}
