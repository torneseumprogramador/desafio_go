package main

import (
	"errors"
	"fmt"
)

// divide realiza a divisão de dois números e retorna um erro se o divisor for zero.
func divide(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("divisão por zero é inválida")
	}
	return dividendo / divisor, nil
}

func main() {
	resultado, err := divide(10, 0) // Tentativa de divisão por zero
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	resultado, _ = divide(10, 2) // Divisão válida
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
