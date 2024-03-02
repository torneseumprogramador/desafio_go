package operacoes

import "testing"

func TestSoma(t *testing.T) {
	result := Soma(1, 2)
	if result != 3 {
		t.Errorf("Soma(1, 2) deu o resultado de : %d; mas deveria ser: 3", result)
	}
}
