package config

import (
	"os"
	"testing"
)

func TestGetEnvComValorPadrao(t *testing.T) {
	result := getEnv("XXX", "um valor")
	if result != "um valor" {
		t.Errorf("A função do getEnv não está retornando valor padrão")
	}
}

func TestGetEnvComVariavelHomeDoSO(t *testing.T) {
	result := getEnv("HOME", "um valor")
	if result == "um valor" {
		t.Errorf("A função do getEnv não está pegando dados da variável de sistema")
	}
}

func TestGetDB(t *testing.T) {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_HOST", "127.0.0.1:3306")
	os.Setenv("DB_NAME", "desafio_go")

	_, erro := GetDB()
	if erro != nil {
		t.Errorf("A conexão falhou: {%s}", erro.Error())
	}
}

func TestGetDBInexistente(t *testing.T) {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_HOST", "127.0.0.1:3306")
	os.Setenv("DB_NAME", "banco_de_dados_inexistente") // Define a variável de ambiente para o teste

	_, erro := GetDB()
	if erro == nil {
		t.Errorf("A conexão com banco de dados 'banco_de_dados_inexistente' deveria falhar")
	}
}

func TestGetDBShowTables(t *testing.T) {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_HOST", "127.0.0.1:3306")
	os.Setenv("DB_NAME", "desafio_go")

	db, erro := GetDB()
	if erro != nil {
		t.Errorf("A conexão falhou: {%s}", erro.Error())
	}

	rows, err := db.Query("show tables")

	if err != nil {
		t.Errorf("Erro no comando SQL: {%s}", err.Error())
	}
	defer rows.Close()
}
