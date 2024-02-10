package modulos

import (
	"aula_go/crud_pet_em_json/modulos/display"
	"fmt"
)

func Menu() {
	for {
		display.LimparTela()

		fmt.Println("=== Escolha um dos itens abaixo: === ")
		fmt.Println("1 - Cadastrar um Pet")
		fmt.Println("2 - Alterar um Pet")
		fmt.Println("3 - Excluir os Pets")
		fmt.Println("4 - Listar os Pets")
		fmt.Println("0 - Sair")

		opcao := display.LerDados()
		sair := false

		switch opcao {
		case "1":
			IncluirPet()
		case "2":
			AlterarPet()
		case "3":
			ExcluirPet()
		case "4":
			ListarPet()
		case "0":
			sair = true
		default:
			fmt.Println("Opção inválida")
			display.Espera(1)
		}

		if sair {
			break
		}
	}
}
