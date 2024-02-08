package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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

type Produto struct {
	Codigo    int
	Nome      string
	Descricao string
	Preco     float32
}

func readInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(input)
}

func programaConsoleProduto() {
	var produtos []Produto
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
			codigo := len(produtos) + 1 // auto increment

			fmt.Println("Digite o nome do produto")
			nomeStr := readInput(reader)

			fmt.Println("Digite a descrição do produto")
			descricaoStr := readInput(reader)

			fmt.Println("Digite o preço do produto")
			precoStr := readInput(reader)
			preco, _ := strconv.ParseFloat(precoStr, 32)

			var produto Produto

			produto.Codigo = codigo
			produto.Nome = nomeStr
			produto.Descricao = descricaoStr
			produto.Preco = float32(preco)

			produtos = append(produtos, produto)
		case 2:
			for _, produto := range produtos {
				fmt.Printf("Código: %d\n", produto.Codigo)
				fmt.Printf("Nome: %s\n", produto.Nome)
				fmt.Printf("Descrição: %s\n", produto.Descricao)
				fmt.Printf("Preço: %.2f\n", produto.Preco)
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
}

func main() {
	programaConsoleProduto()

	// produto := Produto{
	// 	Nome:      "Celular Ix445",
	// 	Descricao: "Um celular bem legal",
	// }

	// mensagem := Apresentar(produto)

	// fmt.Printf("%s", mensagem)
}

func Apresentar(p Produto) string {
	return fmt.Sprintf("Nome do produto \"%s\", Preço do produto \"%s\"", p.Nome, p.Descricao)
}
