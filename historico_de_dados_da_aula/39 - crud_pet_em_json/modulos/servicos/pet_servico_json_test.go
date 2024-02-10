package servicos

import (
	"aula_go/crud_pet_em_json/modulos/enums"
	"aula_go/crud_pet_em_json/modulos/models"
	"os"
	"reflect"
	"testing"
)

func limparJson() {
	os.Remove(CAMINHO_JSON_PATS)
}

func TestAdicionarJson(t *testing.T) {
	limparJson()

	pet := models.Pet{
		Id:   1,
		Nome: "Rex",
		Dono: "João",
		Tipo: enums.Cachorro,
	}

	err := AdicionarJson(pet)
	if err != nil {
		t.Errorf("Erro ao adicionar pet: %v", err)
	}

	pets, err := ListaDePetsJson()
	if err != nil {
		t.Errorf("Erro ao listar pets: %v", err)
	}

	if len(pets) != 1 || !reflect.DeepEqual(pets[0], pet) {
		t.Errorf("Pet adicionado não corresponde ao esperado. Obtido: %+v, Esperado: %+v", pets[0], pet)
	}
}

func listaCom2PetsJson() []models.Pet {
	return []models.Pet{
		models.Pet{
			Id:   1,
			Nome: "Lili",
			Dono: "Maria",
			Tipo: enums.Gato,
		},
		models.Pet{
			Id:   2,
			Nome: "Josevaldo",
			Dono: "Geraldo",
			Tipo: enums.Cachorro,
		},
	}
}

func TestAlterarJson(t *testing.T) {
	limparJson()
	SalvarPets(listaCom2PetsJson())

	pet := models.Pet{
		Id:   1,
		Nome: "Rex",
		Dono: "João",
		Tipo: enums.Cachorro,
	}

	err := AlterarJson(pet)
	if err != nil {
		t.Errorf("Erro ao alterar pet: %v", err)
	}

	petAlterado := BuscarPorIdJson(pet.Id)
	if petAlterado == nil {
		t.Errorf("Erro não encontrei o pet alterado")
	}

	if petAlterado.Nome != pet.Nome {
		t.Errorf("Fiz alteração do nome do pet e o mesmo não alterou")
	}
}

func TestExcluirJson(t *testing.T) {
	limparJson()
	SalvarPets(listaCom2PetsJson())
	err := ExcluirJson(1)
	if err != nil {
		t.Errorf("Erro ao alterar pet: %v", err)
	}

	petsNovo, _ := ListaDePetsJson()
	if len(petsNovo) == 2 {
		t.Errorf("Fiz exclusão do pet e o mesmo não excluiu")
	}
}
