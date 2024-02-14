package servicos

import (
	"errors"
	"http_gin/src/models"
	"strings"
)

func ValidaCampos(pet *models.Pet) error {
	if pet.Id == "" {
		return errors.New("O ID de identificação, não pode ser vazio")
	}

	if strings.TrimSpace(pet.Nome) == "" {
		return errors.New("O nome do pet é obrigatório")
	}

	if strings.TrimSpace(pet.Dono) == "" {
		return errors.New("O dono do pet é obrigatório")
	}

	return nil
}
