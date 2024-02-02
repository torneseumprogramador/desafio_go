package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 10 { // while
	// 	fmt.Println(i)
	// 	i++
	// }

	// i := 0
	// for i < 10 { // while
	// 	fmt.Println(i)
	// 	break
	// }

	// i := 0
	// for i < 10 { // while
	// 	if i == 5 {
	// 		i++
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	for {
		fmt.Println("Escolhe uma das opções abaixo:")
		fmt.Println("1) Mostrar um oi na tela")
		fmt.Println("2) Encerrar")
		var opcao int
		fmt.Scanln(&opcao)

		if opcao == 1 {
			fmt.Println("Oiiiiiiiii")
		} else if opcao == 2 {
			break
		} else {
			fmt.Println("Opção inválida")
		}
	}
}
