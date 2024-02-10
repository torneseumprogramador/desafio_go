package models

import "aula_go/crud_pet_em_json/modulos/enums"

type Pet struct {
	Id   int
	Nome string
	Dono string
	Tipo enums.Tipo
}
