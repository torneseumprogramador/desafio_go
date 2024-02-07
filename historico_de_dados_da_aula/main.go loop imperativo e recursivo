////// código programação funcional ////
// package main

// import "fmt"

// // Função de alta ordem que retorna outra função
// func multiplicador(factor int) func(int) int {
// 	return func(input int) int {
// 		return input * factor
// 	}
// }

// func main() {
// 	dobro := multiplicador(2)
// 	fmt.Println(dobro(5)) // Saída: 10
// }

////// código programação Imperativa ////
// package main

// import "fmt"

// func main() {
// 	sum := 0
// 	for i := 1; i <= 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum) // Saída: 55
// }

////// código programação Procedural ////
// package main

// import "fmt"

// // Uma função que calcula a soma de dois números
// func soma(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	resultado := soma(5, 7)
// 	fmt.Println(resultado) // Saída: 12
// }

////// código programação Orientado a Objetos ////
// package main

// import "fmt"

// // Definição de uma 'struct' para simular uma classe
// type Retangulo struct {
// 	largura, altura float64
// }

// // Método associado à struct Retangulo para calcular a área
// func (r Retangulo) area() float64 {
// 	return r.largura * r.altura
// }

// func main() {
// 	r := Retangulo{largura: 10, altura: 5}
// 	fmt.Println(r.area()) // Saída: 50
// }

// // fazer concorrencia gorotines ///
package main

import "fmt"

func main() {
	fmt.Println("-----------Imperativa-----------")
	sum := 0
	final := 10
	for i := 1; i <= final; i++ {
		sum += i
	}
	fmt.Println(sum) // Saída: 55

	sum2 := 0
	final2 := 2
	for i := 1; i <= final2; i++ {
		sum2 += i
	}
	fmt.Println(sum2) // Saída: 3

	fmt.Println("-----------recursividade-----------")
	// go rotines a possibilidade de deixar paralelo e organizado
	fmt.Println(calc(1, 0, 10)) // Saída: 55
	fmt.Println(calc(1, 0, 2))  // Saída: 3
}

func calc(i int, soma int, final int) int {
	if i > final {
		return soma
	}
	return calc(i+1, soma+i, final)
}
