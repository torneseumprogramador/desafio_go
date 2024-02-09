package modulos

import (
	"aula_go/crud_pet_em_memoria/modulos/display"
	"aula_go/crud_pet_em_memoria/modulos/enums"
	"aula_go/crud_pet_em_memoria/modulos/models"
	"aula_go/crud_pet_em_memoria/modulos/servicos"
	"fmt"
	"strconv"
	"strings"
)

func IncluirPet() {
	display.LimparTela()

	pet := models.Pet{}
	pet.Id = len(models.ListaDePet) + 1
	capturaNomeDonoTipo(&pet)

	// models.ListaDePet = append(models.ListaDePet, pet)
	erro := servicos.Adicionar(&models.ListaDePet, pet)
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

	petCadastrado, _ := localizaPetPorId(id)

	if petCadastrado == nil {
		fmt.Println("Pet não encontrado.")
		display.Espera(1)
		return
	}

	fmt.Println("Dados do Pet localizado:")
	mostrarPet(petCadastrado)
	fmt.Println(strings.Repeat("-", 40))

	capturaNomeDonoTipo(petCadastrado)

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

	pet, i := localizaPetPorId(id)

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
		models.ListaDePet = append(models.ListaDePet[:i], models.ListaDePet[i+1:]...)
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

func localizaPetPorId(id int) (*models.Pet, int) {
	for i, pet := range models.ListaDePet {
		if pet.Id == id {
			return &models.ListaDePet[i], i
		}
	}

	return nil, -1
}

func ListarPet() {
	display.LimparTela()

	if len(models.ListaDePet) == 0 {
		fmt.Println("Não tenho Pets cadastrados")
		display.Espera(2)
		return
	}

	fmt.Println(strings.Repeat("-", 40))

	for _, pet := range models.ListaDePet {
		mostrarPet(&pet)
		fmt.Println(strings.Repeat("-", 40))
	}

	fmt.Println("Digite enter para sair ...")
	display.LerDados()
}
