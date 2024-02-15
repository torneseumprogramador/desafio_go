package servicos

import (
	"encoding/json"
	"errors"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"os"
	"strings"

	"github.com/google/uuid"
)

const CAMINHO_JSON_PATS = "db/pets.json"

type PetServico struct{}

func (ps *PetServico) Lista() ([]models.Pet, error) {
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

func (ps *PetServico) ListaPetView(ds DonoServico) ([]model_views.PetView, error) {
	var pets []models.Pet
	var pets_views []model_views.PetView

	bytes, err := os.ReadFile(CAMINHO_JSON_PATS)
	if err != nil {
		if os.IsNotExist(err) {
			return []model_views.PetView{}, nil // Retorna slice vazia se o arquivo não existir
		}
		return nil, err
	}

	err = json.Unmarshal(bytes, &pets)
	if err != nil {
		return nil, err
	}

	for _, pet := range pets {

		petView := model_views.PetView{}
		petView.Id = pet.Id
		petView.Nome = pet.Nome
		petView.DonoId = pet.DonoId
		petView.Dono = ds.BuscarPorId(pet.DonoId).Nome

		pets_views = append(pets_views, petView)
	}
	return pets_views, nil
}

func (ps *PetServico) BuscarPorId(id string) *models.Pet {
	pets, _ := ps.Lista()
	pet, _ := ps.buscarPorIdStruct(pets, id)
	return pet
}

func (ps *PetServico) Salvar(pets []models.Pet) error {
	bytes, err := json.Marshal(pets)
	if err != nil {
		return err
	}

	return os.WriteFile(CAMINHO_JSON_PATS, bytes, 0644)
}

func (ps *PetServico) Adicionar(pet models.Pet) error {
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

func (ps *PetServico) Alterar(pet models.Pet) error {
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
	pets[index].DonoId = pet.DonoId
	pets[index].Tipo = pet.Tipo

	return ps.Salvar(pets)
}

func (ps *PetServico) Excluir(id string) error {
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

func (ps *PetServico) buscarPorId(pets []models.Pet, id string) (int, error) {
	for i, pet := range pets {
		if pet.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("Pet não encontrado")
}

func (ps *PetServico) BuscarPorDonoId(id string) []models.Pet {
	pets := []models.Pet{}

	lista, _ := ps.Lista()

	for _, pet := range lista {
		if pet.Id == id {
			pets = append(pets, pet)
		}
	}

	return pets
}

func (ps *PetServico) buscarPorIdStruct(pets []models.Pet, id string) (*models.Pet, int) {
	for i, pet := range pets {
		if pet.Id == id {
			return &pet, i
		}
	}

	return nil, -1
}

func (ps *PetServico) validaCampos(pet *models.Pet) error {
	if pet.Id == "" {
		return errors.New("O ID de identificação, não pode ser vazio")
	}

	if strings.TrimSpace(pet.Nome) == "" {
		return errors.New("O nome do pet é obrigatório")
	}

	if strings.TrimSpace(pet.DonoId) == "" {
		return errors.New("O dono do pet é obrigatório")
	}

	return nil
}
