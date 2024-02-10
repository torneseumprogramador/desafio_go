package modulos

import (
	"aula_go/crud_pet_em_json/modulos/display"
	"aula_go/crud_pet_em_json/modulos/enums"
	"aula_go/crud_pet_em_json/modulos/models"
	"aula_go/crud_pet_em_json/modulos/servicos"
	"fmt"
	"strconv"
	"strings"
)

func IncluirPet() {
	display.LimparTela()

	pet := models.Pet{}
	capturaNomeDonoTipo(&pet)

	erro := servicos.AdicionarJson(pet)
	if erro != nil {
		display.LimparTela()
		fmt.Println(erro.Error())
		fmt.Println("Vamos tentar novamente ?")
		display.Espera(2)
		IncluirPet()
	}
}

func capturaNomeDonoTipo(pet *models.Pet) {
	fmt.Println("Digite o nome do Pet")
	pet.Nome = display.LerDados()

	fmt.Println("Digite o nome do Dono do Pet")
	pet.Dono = display.LerDados()

	fmt.Println("Tipo do Pet: \n1 - para cachorro(padrão)\n2 - para gato\n3 - para outros")
	tipoInt, _ := strconv.Atoi(display.LerDados())

	if tipoInt == 2 {
		pet.Tipo = enums.Gato
	} else if tipoInt == 3 {
		pet.Tipo = enums.Outros
	} else {
		pet.Tipo = enums.Cachorro
	}
}

func AlterarPet() {
	display.LimparTela()
	fmt.Println("Digite o Id do Pet")
	idStr := display.LerDados()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Id inválido")
		display.Espera(1)
		AlterarPet()
		return
	}

	petCadastrado := servicos.BuscarPorIdJson(id)

	if petCadastrado == nil {
		fmt.Println("Pet não encontrado.")
		display.Espera(1)
		return
	}

	fmt.Println("Dados do Pet localizado:")
	mostrarPet(petCadastrado)
	fmt.Println(strings.Repeat("-", 40))

	var petParaAlterar models.Pet
	capturaNomeDonoTipo(&petParaAlterar)
	petParaAlterar.Id = petCadastrado.Id

	erro := servicos.AlterarJson(petParaAlterar)
	if erro != nil {
		display.LimparTela()
		fmt.Println(erro.Error())
		display.Espera(2)
		return
	}

	display.LimparTela()
	fmt.Println("Alteração de Pet realizada com sucesso.")
	display.Espera(2)
}

func ExcluirPet() {
	display.LimparTela()
	fmt.Println("Digite o Id do Pet")
	idStr := display.LerDados()
	display.LimparTela()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Id inválido")
		display.Espera(2)
		ExcluirPet()
		return
	}

	pet := servicos.BuscarPorIdJson(id)

	if pet == nil {
		fmt.Println("Pet não encontrado.")
		display.Espera(2)
		return
	}

	fmt.Println("Dados do Pet localizado:")
	mostrarPet(pet)
	fmt.Println(strings.Repeat("-", 40))

	fmt.Println("Deseja realmente excluir ?\ns - Sim\nn - para Não")
	opcao := display.LerDados()
	if opcao == "s" {
		erro := servicos.ExcluirJson(pet.Id)
		if erro != nil {
			display.LimparTela()
			fmt.Println(erro.Error())
			display.Espera(2)
			return
		}
		display.LimparTela()
		fmt.Println("Pet excluído com sucesso.")
		display.Espera(2)
		return
	}
}

func mostrarPet(pet *models.Pet) {
	fmt.Println("Id: ", pet.Id)
	fmt.Println("Nome: ", pet.Nome)
	fmt.Println("Dono: ", pet.Dono)

	if pet.Tipo == enums.Cachorro {
		fmt.Println("Tipo: Cachorro")
	} else if pet.Tipo == enums.Gato {
		fmt.Println("Tipo: Gato")
	} else {
		fmt.Println("Tipo: Outros")
	}
}

func ListarPet() {
	display.LimparTela()

	pets, _ := servicos.ListaDePetsJson()

	if len(pets) == 0 {
		fmt.Println("Não tenho Pets cadastrados")
		display.Espera(2)
		return
	}

	fmt.Println(strings.Repeat("-", 40))

	for _, pet := range pets {
		mostrarPet(&pet)
		fmt.Println(strings.Repeat("-", 40))
	}

	fmt.Println("Digite enter para sair ...")
	display.LerDados()
}
