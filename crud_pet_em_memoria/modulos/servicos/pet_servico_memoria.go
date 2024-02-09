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
