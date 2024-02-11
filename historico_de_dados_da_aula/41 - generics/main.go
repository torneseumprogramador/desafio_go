///////=========== não genérico =============////////

// package main

// import "fmt"

// // MaxInt é uma função específica para inteiros.
// func MaxInt(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // MaxString é uma função específica para strings.
// func MaxString(a, b string) string {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func MaxFloat(a, b float32) float32 {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	fmt.Println(MaxInt(1, 2))        // Output: 2
// 	fmt.Println(MaxFloat(1.5, 2.7))  // Output: 2.7
// 	fmt.Println(MaxString("a", "b")) // Output: "b"
// }

///////=========== genérico =============////////

// package main

// import (
// 	"fmt"
// )

// // Comparable é uma interface que define o método CompareTo.
// // O método CompareTo deve retornar um valor negativo se o primeiro valor for menor que o segundo,
// // zero se forem iguais, e um valor positivo se o primeiro valor for maior.
// type Comparable[T any] interface {
// 	CompareTo(T) int
// }

// // Max retorna o maior de dois valores de um tipo que implementa a interface Comparable.
// func Max[T Comparable[T]](a, b T) T {
// 	if a.CompareTo(b) > 0 {
// 		return a
// 	}
// 	return b
// }

// type Int int

// // CompareTo compara dois inteiros.
// func (a Int) CompareTo(b Int) int {
// 	return int(a - b)
// }

// type String string

// func (a String) CompareTo(b String) int {
// 	if a == b {
// 		return 0
// 	} else if a > b {
// 		return 1
// 	} else {
// 		return -1
// 	}
// }

// type Float float32

// func (a Float) CompareTo(b Float) int {
// 	if a == b {
// 		return 0
// 	} else if a > b {
// 		return 1
// 	} else {
// 		return -1
// 	}
// }

// func main() {
// 	fmt.Println(Max(Int(1), Int(2)))           // Output: 2
// 	fmt.Println(Max(String("a"), String("b"))) // Output: b
// 	fmt.Println(Max(Float(1.4), Float(1.2)))   // Output: 1.4
// }

// package main

// import (
// 	"fmt"
// )

// // PrintGeneric é uma função genérica que pode imprimir qualquer tipo de dado.
// func PrintGeneric[T any](value T) {
// 	fmt.Println(value)
// }

// func main() {
// 	PrintGeneric("Hello, world!") // Imprime uma string
// 	PrintGeneric(123)             // Imprime um inteiro
// 	PrintGeneric(45.67)           // Imprime um float
// 	PrintGeneric([]int{1, 2, 3})  // Imprime uma slice de inteiros
// 	PrintGeneric(true)            // Imprime uma slice de inteiros
// }

// /// ====== calcular salario =======

package main

import (
	"fmt"
)

type ContratoCalculoImposto interface {
	CalculaImposto() float32
}

type Estado int

const (
	SaoPaulo Estado = iota
	RioDeJaneiro
	MinasGerais
	Natal
	Maceio
	Fortaleza
	Amazonas
)

// PrintGeneric é uma função genérica que pode imprimir qualquer tipo de dado.
func CalcularSalario[T ContratoCalculoImposto](value T, estado Estado) float32 {
	r := value.CalculaImposto()
	if estado == SaoPaulo {
		r += 100
	} else if estado == RioDeJaneiro {
		r += 50
	} else if estado == MinasGerais {
		r += 36
	}

	return r
}

// se passei um int = - 10%, dependendo do estado fazemos algo diferente depois
type MeuInt int

func (mi MeuInt) CalculaImposto() float32 {
	return float32(mi) * 10 / 100
}

// se passei um float = - 15%, dependendo do estado fazemos algo diferente depois
type MeuFloat float32

func (mf MeuFloat) CalculaImposto() float32 {
	return float32(mf) * 15 / 100
}

func main() {
	r1 := CalcularSalario(MeuInt(1000), SaoPaulo)
	r2 := CalcularSalario(MeuFloat(1000.00), RioDeJaneiro)

	fmt.Println("Int: ", r1)
	fmt.Println("Float: ", r2)
}
