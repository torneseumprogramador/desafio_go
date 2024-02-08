package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simula a obtenção de dados da API. Retorna um erro aleatoriamente para simular falhas.
func getDataFromAPI(page int) ([]string, error) {
	time.Sleep(1 * time.Second) // Simula a latência da conexão HTTP
	// Simula um erro aleatório
	if rand.Intn(10) < 3 { // 30% de chance de erro
		return nil, errors.New("falha ao obter dados da API")
	}
	// Simula dados retornados com sucesso
	return []string{"data1", "data2"}, nil
}

// Simula a inserção de dados no banco de dados.
func insertDataToDB(data []string) error {
	time.Sleep(500 * time.Microsecond) // Simula a latência de rede ao salvar no banco

	// Simula um erro aleatório
	if rand.Intn(10) < 2 { // 20% de chance de erro
		return errors.New("falha ao inserir dados no banco de dados")
	}
	// Simula uma inserção bem-sucedida
	return nil
}

func processPages(startPage, endPage int, wg *sync.WaitGroup) {
	defer wg.Done() // Garante que o WaitGroup seja notificado quando a goroutine terminar
	for page := startPage; page <= endPage; page++ {
		fmt.Printf("Processando página %d\n", page)
		data, err := getDataFromAPI(page)
		if err != nil {
			fmt.Printf("Erro ao obter dados da página %d: %s\n", page, err)
			continue // Pula para a próxima iteração do loop
		}

		err = insertDataToDB(data)
		if err != nil {
			fmt.Printf("Erro ao inserir dados da página %d no banco de dados: %s\n", page, err)
			continue
		}

		fmt.Printf("Página %d processada com sucesso\n", page)
	}
}

func main() {
	var wg sync.WaitGroup

	// Divide o trabalho em dois blocos de páginas e processa cada bloco em sua própria goroutine
	wg.Add(1)
	go processPages(1, 10, &wg)

	wg.Add(1)
	go processPages(10, 20, &wg)

	wg.Add(1)
	go processPages(20, 30, &wg)

	wg.Wait() // Aguarda todas as goroutines terminarem
	fmt.Println("Migração concluída.")
}
