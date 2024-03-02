package models

import (
	"fmt"
	"testing"
)

func TestInstancia(t *testing.T) {
	adm := Administrador{}
	adm.Id = "teste"
	adm.Nome = "teste"
	adm.Email = "teste"
	adm.Senha = "teste"

	str := fmt.Sprintf("%+v", adm)
	comoDeveFicar := "{tableName:{} Id:teste Nome:teste Email:teste Senha:teste Super:false}"
	if str != comoDeveFicar {
		t.Errorf("Erro: deveria vir assim: %s e venho assim %s", comoDeveFicar, str)
	}
}

func TestInstanciaOutraForma(t *testing.T) {
	adm := Administrador{
		Id:    "teste",
		Nome:  "teste",
		Email: "teste",
		Senha: "teste",
	}

	str := fmt.Sprintf("%+v", adm)
	comoDeveFicar := "{tableName:{} Id:teste Nome:teste Email:teste Senha:teste Super:false}"
	if str != comoDeveFicar {
		t.Errorf("Erro: deveria vir assim: %s e venho assim %s", comoDeveFicar, str)
	}
}
