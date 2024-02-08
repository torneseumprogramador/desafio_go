package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Simula a obtenção de dados da API. Retorna um erro aleatoriamente para simular falhas.
func getDataFromAPI(page int) ([]string, error) {
	time.Sleep(1 * time.Second)
	// Simula um erro aleatório
	probabilidade := rand.Intn(10)
	if probabilidade < 3 { // 30% de chance de erro
		return nil, errors.New("falha ao obter dados da API")
	}
	// Simula dados retornados com sucesso
	return []string{"data1", "data2"}, nil
}

// Simula a inserção de dados no banco de dados.
func insertDataToDB(data []string) error {
	time.Sleep(500 * time.Microsecond)
	// Simula um erro aleatório
	probabilidade := rand.Intn(10)
	if probabilidade < 2 { // 20% de chance de erro
		return errors.New("falha ao inserir dados no banco de dados")
	}
	// Simula uma inserção bem-sucedida
	return nil
}

func main() {
	totalPages := 30 // Simula 5 páginas de dados para migrar

	for page := 1; page <= totalPages; page++ {
		fmt.Printf("Processando página %d\n", page)
		data, err := getDataFromAPI(page)
		if err != nil {
			fmt.Printf("Erro ao obter dados da página %d: %s\n", page, err)
			continue // Pula para a próxima iteração do loop
		}

		err = insertDataToDB(data)
		if err != nil {
			fmt.Printf("Erro ao inserir dados da página %d no banco de dados: %s\n", page, err)
			// Decide continuar com o próximo registro, mesmo após o erro
			continue
		}

		fmt.Printf("Página %d processada com sucesso\n", page)
	}
	fmt.Println("Migração concluída.")
}
