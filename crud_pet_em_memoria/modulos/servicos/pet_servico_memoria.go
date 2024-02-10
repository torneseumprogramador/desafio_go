package servicos

import (
	"aula_go/crud_pet_em_memoria/modulos/models"
	"errors"
	"strings"
)

func Adicionar(pets *[]models.Pet, pet models.Pet) error {
	erro := validaCampos(&pet)
	if erro != nil {
		return erro
	}

	*pets = append(*pets, pet)

	return nil
}

func Alterar(pets *[]models.Pet, pet models.Pet) error {
	petDoDB, _ := BuscarPorId(*pets, pet.Id)
	if petDoDB == nil {
		return errors.New("Pet não encontrado")
	}

	erro := validaCampos(&pet)
	if erro != nil {
		return erro
	}

	petDoDB.Nome = pet.Nome
	petDoDB.Dono = pet.Dono
	petDoDB.Tipo = pet.Tipo

	return nil
}

func Excluir(pets []models.Pet, id int) ([]models.Pet, error) {
	pet, i := BuscarPorId(pets, id)
	if pet == nil {
		return pets, errors.New("Pet não encontrado")
	}

	pets = append(pets[:i], pets[i+1:]...)
	return pets, nil
}

func BuscarPorId(pets []models.Pet, id int) (*models.Pet, int) {
	for i, pet := range pets {
		if pet.Id == id {
			petEncontrado := &pets[i]
			return petEncontrado, i
		}
	}

	return nil, -1
}

func validaCampos(pet *models.Pet) error {
	if pet.Id <= 0 {
		return errors.New("O ID de identificação, não pode ser <= 0")
	}

	if strings.TrimSpace(pet.Nome) == "" {
		return errors.New("O nome do pet é obrigatório")
	}

	if strings.TrimSpace(pet.Dono) == "" {
		return errors.New("O dono do pet é obrigatório")
	}

	return nil
}
