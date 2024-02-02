package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	/*
		Sr Roberdo é um dono de um brechó e precisa controlar o estoque
		dos seus produtos.

		faça um programa onde é possivel fazer o cadastro em memória
		dos produtos do brechó do Roberto.

		Fluxo:

		Menu: O que vc deseja fazer ?
		1 - Cadastrar produto
		2 - Listar produtos cadastrados
		3 - Sair do programa

		os produtos terão os seguintes dados
		codigo, nome, descricao, preco

		Recusos a serem utilizados:
		if, for, array fixo, para colunas, slice para linhas dos produtos
	*/

	var produtos [][4]string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("O que você deseja fazer?")
		fmt.Println("1 - Cadastrar produto")
		fmt.Println("2 - Listar produtos cadastrados")
		fmt.Println("3 - Sair do programa")

		fmt.Println("Escolha uma opção: ")
		opcaoStr, _ := reader.ReadString('\n')
		opcaoStr = strings.TrimSpace(opcaoStr) // Remove a newline character from the input
		var opcao int
		fmt.Sscanf(opcaoStr, "%d", &opcao)

		clearScreen()

		switch opcao {
		case 1:
			var produto [4]string

			fmt.Println("Digite o código do produto")
			produto[0], _ = reader.ReadString('\n')
			produto[0] = strings.TrimSpace(produto[0])

			fmt.Println("Digite o nome do produto")
			produto[1], _ = reader.ReadString('\n')
			produto[1] = strings.TrimSpace(produto[1])

			fmt.Println("Digite a descrição do produto")
			produto[2], _ = reader.ReadString('\n')
			produto[2] = strings.TrimSpace(produto[2])

			fmt.Println("Digite o preço do produto")
			produto[3], _ = reader.ReadString('\n')
			produto[3] = strings.TrimSpace(produto[3])

			produtos = append(produtos, produto)
		case 2:
			for _, produto := range produtos {
				fmt.Printf("Código: %s\n", produto[0])
				fmt.Printf("Nome: %s\n", produto[1])
				fmt.Printf("Descrição: %s\n", produto[2])
				fmt.Printf("Preço: %s\n", produto[3])
				fmt.Println("-----------------------------")
			}

			// time.Sleep(5 * time.Second)
			fmt.Println("Digite o enter para continuar ...")
			fmt.Scanln()
		case 3:
			fmt.Println("Saindo do programa ...")
			return
		default:
			fmt.Println("Opção inválida")
		}

		clearScreen()
	}

	// var meuSlice []int // slice = array dinamico
	// fmt.Printf("A quantidade é %d\n", len(meuSlice))

	// meuSlice = append(meuSlice, 10) // adicionando 1 item no meu slice (array dinamico)
	// fmt.Printf("A quantidade é %d\n", len(meuSlice))

	// // fazer pop dos dados do array dinamico (Slice)
	// meuSlice = meuSlice[:len(meuSlice)-1]
	// fmt.Printf("A quantidade é %d\n", len(meuSlice))

	/*















		// depois de resolver o exercicio, falar sobre ponteiros, funções e separação de modulos

		i := 42

		var p *int // ponteiro definido
		p = &i     // passagem da referencia de i

		// i = 7 // no Rust isso não é possivel por não ter garbage colector

		fmt.Printf("Valor de i (%d) o endereço de memória (%p)\n", i, &p)
		fmt.Printf("Valor de p (%d) o endereço de memória (%p)\n", *p, p)

		fmt.Println("======== valor por copia ============")

		x := 42
		z := x

		fmt.Printf("Valor de x (%d) o endereço de memória (%p)\n", x, &x)
		fmt.Printf("Valor de z (%d) o endereço de memória (%p)\n", z, &z)

	*/
}
