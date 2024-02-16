package servicos

import (
	"database/sql"
	"errors"
	"http_gin/src/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type AdministradorServicoMySql struct {
	DB *sql.DB
}

// Lista todos os administradores
func (as *AdministradorServicoMySql) Lista() ([]models.Administrador, error) {
	var administradores []models.Administrador

	rows, err := as.DB.Query("SELECT id, nome, email, senha FROM administradores")
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

func (as *AdministradorServicoMySql) BuscarPorId(id string) (*models.Administrador, error) {
	var adm models.Administrador

	// Prepara a consulta SQL para buscar o adm pelo ID
	query := "SELECT id, nome, email, senha FROM administradores WHERE id = ?"
	err := as.DB.QueryRow(query, id).Scan(&adm.Id, &adm.Nome, &adm.Email, &adm.Senha)

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
func (as *AdministradorServicoMySql) Adicionar(adm models.Administrador) error {
	if adm.Id == "" {
		adm.Id = uuid.New().String()
	}

	erro := as.validaCampos(&adm)
	if erro != nil {
		return erro
	}

	_, err := as.DB.Exec("INSERT INTO administradores (id, nome, email, senha) VALUES (?, ?, ?, ?)",
		adm.Id, adm.Nome, adm.Email, adm.Senha)

	return err
}

// Altera um administrador existente
func (as *AdministradorServicoMySql) Alterar(adm models.Administrador) error {
	erro := as.validaCampos(&adm)
	if erro != nil {
		return erro
	}

	_, err := as.DB.Exec("UPDATE administradores SET nome = ?, email = ?, senha = ? WHERE id = ?",
		adm.Nome, adm.Email, adm.Senha, adm.Id)

	return err
}

// Exclui um administrador pelo ID
func (as *AdministradorServicoMySql) Excluir(id string) error {
	_, err := as.DB.Exec("DELETE FROM administradores WHERE id = ?", id)
	return err
}

func (as *AdministradorServicoMySql) validaCampos(pet *models.Administrador) error {
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
