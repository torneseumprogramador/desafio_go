package main

import "fmt"

func main() {
	// var meuArray [5]int
	// meuArray[0] = 1
	// meuArray[1] = 2
	// meuArray[2] = 3
	// meuArray[3] = 4
	// meuArray[4] = 5

	// var meuArray = [5]int{10, 20, 30, 40, 50}

	// fmt.Println(meuArray[2])

	// for i, v := range meuArray {
	// 	fmt.Println("Índice:", i, "Valor:", v)
	// }

	// matriz é uma tabela de dados em memória

	// var matriz [2][3]int

	// matriz[0][0] = 11 // coluna 1 da linha 0 da matriz
	// matriz[0][1] = 14 // coluna 2 da linha 0 da matriz
	// matriz[0][2] = 16 // coluna 3 da linha 0 da matriz

	// matriz[1][0] = 21 // coluna 1 da linha 2 da matriz
	// matriz[1][1] = 24 // coluna 2 da linha 2 da matriz
	// matriz[1][2] = 26 // coluna 3 da linha 2 da matriz

	matriz := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for iLinha, vLinha := range matriz {
		// for das linhas
		fmt.Println("Linhas - Índice:", iLinha, "Valor:", vLinha)

		for iColuna, vColuna := range matriz[iLinha] {
			// for das colunas
			fmt.Println("Colunas - Índice:", iColuna, "Valor:", vColuna)
		}
	}

	fmt.Println("=======================")

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
}
