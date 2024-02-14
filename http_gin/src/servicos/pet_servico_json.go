package servicos

import (
	"encoding/json"
	"errors"
	"http_gin/src/models"
	"os"

	"github.com/google/uuid"
)

const CAMINHO_JSON_PATS = "pets.json"

func ListaDePetsJson() ([]models.Pet, error) {
	var pets []models.Pet

	bytes, err := os.ReadFile(CAMINHO_JSON_PATS)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Pet{}, nil // Retorna slice vazia se o arquivo não existir
		}
		return nil, err
	}

	err = json.Unmarshal(bytes, &pets)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func BuscarPorIdJson(id string) *models.Pet {
	pets, _ := ListaDePetsJson()
	pet, _ := buscarPorIdStruct(pets, id)
	return pet
}

// salvarPets no arquivo
func SalvarPets(pets []models.Pet) error {
	bytes, err := json.Marshal(pets)
	if err != nil {
		return err
	}

	return os.WriteFile(CAMINHO_JSON_PATS, bytes, 0644)
}

// Adicionar um pet ao arquivo
func AdicionarJson(pet models.Pet) error {
	pets, err := ListaDePetsJson()
	if err != nil {
		return err
	}

	if pet.Id == "" {
		pet.Id = uuid.New().String()
	}

	erro := ValidaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets = append(pets, pet)
	return SalvarPets(pets)
}

// Alterar um pet no arquivo
func AlterarJson(pet models.Pet) error {
	pets, err := ListaDePetsJson()
	if err != nil {
		return err
	}

	index, err := buscarPorId(pets, pet.Id)
	if err != nil {
		return err
	}

	erro := ValidaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets[index].Nome = pet.Nome
	pets[index].Dono = pet.Dono
	pets[index].Tipo = pet.Tipo

	return SalvarPets(pets)
}

// Excluir um pet do arquivo
func ExcluirJson(id string) error {
	pets, err := ListaDePetsJson()
	if err != nil {
		return err
	}

	index, err := buscarPorId(pets, id)
	if err != nil {
		return err
	}

	pets = append(pets[:index], pets[index+1:]...)
	return SalvarPets(pets)
}

// BuscarPorId busca um pet pelo ID
func buscarPorId(pets []models.Pet, id string) (int, error) {
	for i, pet := range pets {
		if pet.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("Pet não encontrado")
}

// BuscarPorId busca um pet pelo ID
func buscarPorIdStruct(pets []models.Pet, id string) (*models.Pet, int) {
	for i, pet := range pets {
		if pet.Id == id {
			return &pet, i
		}
	}

	return nil, -1
}
