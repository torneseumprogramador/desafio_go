package servicos

import (
	"encoding/json"
	"errors"
	"http_gin/src/models"
	"os"
	"strings"

	"github.com/google/uuid"
)

const CAMINHO_JSON_DONOS = "db/donos.json"

type DonoServico struct{}

func (ps *DonoServico) Lista() ([]models.Dono, error) {
	var pets []models.Dono

	bytes, err := os.ReadFile(CAMINHO_JSON_DONOS)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Dono{}, nil // Retorna slice vazia se o arquivo não existir
		}
		return nil, err
	}

	err = json.Unmarshal(bytes, &pets)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func (ps *DonoServico) BuscarPorId(id string) *models.Dono {
	pets, _ := ps.Lista()
	pet, _ := ps.buscarPorIdStruct(pets, id)
	return pet
}

func (ps *DonoServico) Salvar(pets []models.Dono) error {
	bytes, err := json.Marshal(pets)
	if err != nil {
		return err
	}

	return os.WriteFile(CAMINHO_JSON_DONOS, bytes, 0644)
}

func (ps *DonoServico) Adicionar(pet models.Dono) error {
	pets, err := ps.Lista()
	if err != nil {
		return err
	}

	if pet.Id == "" {
		pet.Id = uuid.New().String()
	}

	erro := ps.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets = append(pets, pet)
	return ps.Salvar(pets)
}

func (ps *DonoServico) Alterar(pet models.Dono) error {
	pets, err := ps.Lista()
	if err != nil {
		return err
	}

	index, err := ps.buscarPorId(pets, pet.Id)
	if err != nil {
		return err
	}

	erro := ps.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets[index].Nome = pet.Nome
	pets[index].Telefone = pet.Telefone

	return ps.Salvar(pets)
}

func (ps *DonoServico) Excluir(id string) error {
	pets, err := ps.Lista()
	if err != nil {
		return err
	}

	index, err := ps.buscarPorId(pets, id)
	if err != nil {
		return err
	}

	pets = append(pets[:index], pets[index+1:]...)
	return ps.Salvar(pets)
}

func (ps *DonoServico) buscarPorId(pets []models.Dono, id string) (int, error) {
	for i, pet := range pets {
		if pet.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("Pet não encontrado")
}

func (ps *DonoServico) buscarPorIdStruct(pets []models.Dono, id string) (*models.Dono, int) {
	for i, pet := range pets {
		if pet.Id == id {
			return &pet, i
		}
	}

	return nil, -1
}

func (ps *DonoServico) validaCampos(pet *models.Dono) error {
	if pet.Id == "" {
		return errors.New("O ID de identificação, não pode ser vazio")
	}

	if strings.TrimSpace(pet.Nome) == "" {
		return errors.New("O nome é obrigatório")
	}

	if strings.TrimSpace(pet.Telefone) == "" {
		return errors.New("O telefone obrigatório")
	}

	return nil
}
