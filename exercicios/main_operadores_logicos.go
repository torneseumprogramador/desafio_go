package main

import "fmt"

// func main() {
// 	faixa_idade_inicial := 23
// 	faixa_idade_final := 35
// 	idade := 0

// 	fmt.Print("Digite sua idade: \n")
// 	fmt.Scanln(&idade)

// 	if idade >= faixa_idade_inicial && idade <= faixa_idade_final {
// 		fmt.Printf("A pessoa tem %d, faz parte do seu publico alvo", idade)
// 	} else {
// 		fmt.Printf("A pessoa tem %d, não faz parte do seu publico alvo", idade)
// 	}
// }

// func main() {
// 	idade := 16
// 	possuiPermissao := true
// 	if idade >= 18 || possuiPermissao {
// 		fmt.Println("Acesso permitido.")
// 	} else {
// 		fmt.Println("Acesso negado.")
// 	}
// }

func main() {
	idade := 18
	if !(idade >= 18) {
		fmt.Println("Acesso permitido.")
	} else {
		fmt.Println("Acesso negado.")
	}
}
