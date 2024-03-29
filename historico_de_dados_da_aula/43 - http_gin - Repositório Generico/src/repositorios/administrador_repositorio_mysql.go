package repositorios

import (
	"database/sql"
	"errors"
	"http_gin/src/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type AdministradorRepositorioMySql struct {
	DB *sql.DB
}

// Lista todos os administradores
func (ar *AdministradorRepositorioMySql) Lista() ([]models.Administrador, error) {
	var administradores []models.Administrador

	rows, err := ar.DB.Query("SELECT id, nome, email, senha FROM administradores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var adm models.Administrador
		if err := rows.Scan(&adm.Id, &adm.Nome, &adm.Email, &adm.Senha); err != nil {
			return nil, err
		}
		administradores = append(administradores, adm)
	}

	return administradores, nil
}

func (ar *AdministradorRepositorioMySql) BuscarPorId(id string) (*models.Administrador, error) {
	var adm models.Administrador

	// Prepara a consulta SQL para buscar o adm pelo ID
	query := "SELECT id, nome, email, senha FROM administradores WHERE id = ?"
	err := ar.DB.QueryRow(query, id).Scan(&adm.Id, &adm.Nome, &adm.Email, &adm.Senha)

	if err != nil {
		if err == sql.ErrNoRows {
			// Nenhum resultado encontrado
			return nil, nil
		}
		// Algum outro erro ocorreu
		return nil, err
	}

	return &adm, nil
}

// Adiciona um novo administrador
func (ar *AdministradorRepositorioMySql) Adicionar(adm models.Administrador) error {
	if adm.Id == "" {
		adm.Id = uuid.New().String()
	}

	erro := ar.validaCampos(&adm)
	if erro != nil {
		return erro
	}

	_, err := ar.DB.Exec("INSERT INTO administradores (id, nome, email, senha) VALUES (?, ?, ?, ?)",
		adm.Id, adm.Nome, adm.Email, adm.Senha)

	return err
}

// Altera um administrador existente
func (ar *AdministradorRepositorioMySql) Alterar(adm models.Administrador) error {
	erro := ar.validaCampos(&adm)
	if erro != nil {
		return erro
	}

	_, err := ar.DB.Exec("UPDATE administradores SET nome = ?, email = ?, senha = ? WHERE id = ?",
		adm.Nome, adm.Email, adm.Senha, adm.Id)

	return err
}

// Exclui um administrador pelo ID
func (ar *AdministradorRepositorioMySql) Excluir(id string) error {
	_, err := ar.DB.Exec("DELETE FROM administradores WHERE id = ?", id)
	return err
}

func (ar *AdministradorRepositorioMySql) validaCampos(pet *models.Administrador) error {
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
