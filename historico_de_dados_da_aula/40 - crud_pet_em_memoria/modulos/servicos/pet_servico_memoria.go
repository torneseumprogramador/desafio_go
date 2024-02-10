package servicos

import (
	"aula_go/crud_pet_em_memoria/modulos/models"
	"errors"
)

func Adicionar(pets *[]models.Pet, pet models.Pet) error {
	erro := ValidaCampos(&pet)
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

	erro := ValidaCampos(&pet)
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
