package servicos

import (
	"encoding/json"
	"errors"
	"http_gin/src/models"
	"os"
	"strings"

	"github.com/google/uuid"
)

const CAMINHO_JSON_ADMINISTRADORES = "db/administradores.json"

type AdministradorServico struct{}

func (ps *AdministradorServico) Lista() ([]models.Administrador, error) {
	var pets []models.Administrador

	bytes, err := os.ReadFile(CAMINHO_JSON_ADMINISTRADORES)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Administrador{}, nil // Retorna slice vazia se o arquivo não existir
		}
		return nil, err
	}

	err = json.Unmarshal(bytes, &pets)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func (ps *AdministradorServico) BuscarPorId(id string) *models.Administrador {
	pets, _ := ps.Lista()
	pet, _ := ps.buscarPorIdStruct(pets, id)
	return pet
}

func (ps *AdministradorServico) Salvar(pets []models.Administrador) error {
	bytes, err := json.Marshal(pets)
	if err != nil {
		return err
	}

	return os.WriteFile(CAMINHO_JSON_ADMINISTRADORES, bytes, 0644)
}

func (ps *AdministradorServico) Adicionar(pet models.Administrador) error {
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

func (ps *AdministradorServico) Alterar(pet models.Administrador) error {
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
	pets[index].Email = pet.Email
	pets[index].Senha = pet.Senha

	return ps.Salvar(pets)
}

func (ps *AdministradorServico) Excluir(id string) error {
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

func (ps *AdministradorServico) buscarPorId(pets []models.Administrador, id string) (int, error) {
	for i, pet := range pets {
		if pet.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("Pet não encontrado")
}

func (ps *AdministradorServico) BuscarPorEmailSenha(email string, senha string) *models.Administrador {
	lista, _ := ps.Lista()
	for _, adm := range lista {
		if adm.Email == email && adm.Senha == senha {
			return &adm
		}
	}

	return nil
}

func (ps *AdministradorServico) buscarPorIdStruct(pets []models.Administrador, id string) (*models.Administrador, int) {
	for i, pet := range pets {
		if pet.Id == id {
			return &pet, i
		}
	}

	return nil, -1
}

func (ps *AdministradorServico) validaCampos(pet *models.Administrador) error {
	if pet.Id == "" {
		return errors.New("O ID de identificação, não pode ser vazio")
	}

	if strings.TrimSpace(pet.Nome) == "" {
		return errors.New("O nome é obrigatório")
	}

	if strings.TrimSpace(pet.Email) == "" {
		return errors.New("O email obrigatório")
	}

	if strings.TrimSpace(pet.Senha) == "" {
		return errors.New("A Senha obrigatória")
	}

	return nil
}
