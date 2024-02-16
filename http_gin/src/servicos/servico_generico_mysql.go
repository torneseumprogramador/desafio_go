package servicos

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type GenericService struct {
	DB    *sql.DB
	Table string
	Type  reflect.Type // Tipo da entidade
}

// Lista implementa a operação de listar entidades.
func (gs *GenericService) Lista() ([]interface{}, error) {
	rows, err := gs.DB.Query(fmt.Sprintf("SELECT * FROM %s", gs.Table))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []interface{}
	for rows.Next() {
		entidade := reflect.New(gs.Type).Interface()
		err := rows.Scan(entidade) // Isso requer ajuste para mapear corretamente os campos
		if err != nil {
			log.Println("Erro ao escanear a entidade:", err)
			continue
		}
		result = append(result, entidade)
	}
	return result, nil
}
