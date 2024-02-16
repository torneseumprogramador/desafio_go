package servicos

import (
	"database/sql"
	"errors"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type PetServico struct {
	DB *sql.DB
}

// Lista todos os pets
func (ps *PetServico) Lista() ([]models.Pet, error) {
	var pets []models.Pet

	rows, err := ps.DB.Query("SELECT id, nome, dono_id, tipo FROM pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pet models.Pet
		if err := rows.Scan(&pet.Id, &pet.Nome, &pet.DonoId, &pet.Tipo); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

func (ps *PetServico) BuscarPorId(id string) (*models.Pet, error) {
	var pet models.Pet

	// Prepara a consulta SQL para buscar o pet pelo ID
	query := "SELECT id, nome, dono_id, tipo FROM pets WHERE id = ?"
	err := ps.DB.QueryRow(query, id).Scan(&pet.Id, &pet.Nome, &pet.DonoId, &pet.Tipo)

	if err != nil {
		if err == sql.ErrNoRows {
			// Nenhum resultado encontrado
			return nil, nil
		}
		// Algum outro erro ocorreu
		return nil, err
	}

	return &pet, nil
}

// Lista pets com informações do dono
func (ps *PetServico) ListaPetView(ds DonoServico) ([]model_views.PetView, error) {
	pets, err := ps.Lista()
	if err != nil {
		return nil, err
	}

	var petViews []model_views.PetView
	for _, pet := range pets {
		dono, err := ds.BuscarPorId(pet.DonoId)
		if err != nil {
			continue // ou tratar erro
		}

		petView := model_views.PetView{
			Id:     pet.Id,
			Nome:   pet.Nome,
			DonoId: pet.DonoId,
			Dono:   dono.Nome, // Assumindo que BuscarPorId retorna o nome do dono
		}

		petViews = append(petViews, petView)
	}

	return petViews, nil
}

// Adiciona um novo pet
func (ps *PetServico) Adicionar(pet models.Pet) error {
	if pet.Id == "" {
		pet.Id = uuid.New().String()
	}

	erro := ps.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	_, err := ps.DB.Exec("INSERT INTO pets (id, nome, dono_id, tipo) VALUES (?, ?, ?, ?)",
		pet.Id, pet.Nome, pet.DonoId, pet.Tipo)

	return err
}

// Altera um pet existente
func (ps *PetServico) Alterar(pet models.Pet) error {
	erro := ps.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	_, err := ps.DB.Exec("UPDATE pets SET nome = ?, dono_id = ?, tipo = ? WHERE id = ?",
		pet.Nome, pet.DonoId, pet.Tipo, pet.Id)

	return err
}

// Exclui um pet pelo ID
func (ps *PetServico) Excluir(id string) error {
	_, err := ps.DB.Exec("DELETE FROM pets WHERE id = ?", id)
	return err
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
