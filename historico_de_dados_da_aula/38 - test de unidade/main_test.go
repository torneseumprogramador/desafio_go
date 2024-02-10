package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var resultado int
	var esperado int = 3

	// Act (Roda a função que irá ser testada)
	resultado = Soma(1, 2)

	// Assert (Verifica o resultado sobre o teste)
	if esperado != resultado {
		t.Errorf("Soma(1, 2) = %d; espero %d", esperado, resultado)
	}
}
