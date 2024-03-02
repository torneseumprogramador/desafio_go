package repositorios

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"tests/config"
	"tests/models"
)

func TestGenericoRepositorioMySql(t *testing.T) {
	repo := GenericoRepositorioMySql[models.Administrador]{}

	str := fmt.Sprintf("%+v", repo)
	comoDeveFicar := "{DB:<nil>}"
	if str != comoDeveFicar {
		t.Errorf("Erro: deveria vir assim: %s e venho assim %s", comoDeveFicar, str)
	}
}

func TestGenericoRepositorioMySqlPassandoConexao(t *testing.T) {
	db, _ := config.GetDB()
	repo := GenericoRepositorioMySql[models.Administrador]{
		DB: db,
	}

	if repo.DB == nil {
		t.Errorf("Conexão indevida")
	}
}

func TestGetTableName(t *testing.T) {
	repo := GenericoRepositorioMySql[models.Administrador]{}

	table := repo.getTableName()

	if table != "administradores" {
		t.Errorf("A tabela deveria vir administradores porém venho %s", table)
	}
}

func TestLista(t *testing.T) {
	db, _ := config.GetDB()
	repo := GenericoRepositorioMySql[models.Administrador]{
		DB: db,
	}

	adms, erro := repo.Lista()

	if erro != nil {
		t.Errorf("Erro ao buscar administrasdores %s", erro.Error())
	}

	esperado := reflect.TypeOf([]models.Administrador{})

	if reflect.TypeOf(adms) != esperado {
		t.Errorf("Tipo esperado: %v, obtido: %v", esperado, reflect.TypeOf(adms))
	}
}

func TestAdicionar(t *testing.T) {
	db, _ := config.GetDB()
	truncateTable(db)

	repo := GenericoRepositorioMySql[models.Administrador]{
		DB: db,
	}

	adm := models.Administrador{
		Id:    "1",
		Nome:  "teste",
		Email: "teste",
		Senha: "teste",
	}

	erroAdd := repo.Adicionar(adm)
	if erroAdd != nil {
		t.Errorf("Erro ao cadastrar administrador %s", erroAdd.Error())
	}
}

func TestBuscaPorId(t *testing.T) {
	db, _ := config.GetDB()
	truncateTable(db)

	repo := GenericoRepositorioMySql[models.Administrador]{
		DB: db,
	}

	adm := models.Administrador{
		Id:    "1",
		Nome:  "teste",
		Email: "teste",
		Senha: "teste",
	}

	erroAdd := repo.Adicionar(adm)
	if erroAdd != nil {
		t.Errorf("Erro ao cadastrar administrador %s", erroAdd.Error())
	}

	admDb, erro := repo.BuscaPorId("1")

	if erro != nil {
		t.Errorf("Erro ao buscar administrasdor por id %s", erro.Error())
	}

	esperado := reflect.TypeOf(models.Administrador{})

	if reflect.TypeOf(*admDb) != esperado {
		t.Errorf("Tipo esperado: %v, obtido: %v", esperado, reflect.TypeOf(admDb))
	}

	if admDb.Id == "" {
		t.Errorf("O administrador deveria existir")
	}
}

func truncateTable(db *sql.DB) {
	rows, err := db.Query("truncate table administradores")
	if err != nil {
		fmt.Printf("Erro no comando SQL: {%s}\n", err.Error())
	}
	defer rows.Close()
}

func TestBuscaPorIdInexistente(t *testing.T) {
	db, _ := config.GetDB()
	truncateTable(db)

	repo := GenericoRepositorioMySql[models.Administrador]{
		DB: db,
	}

	admDb, erro := repo.BuscaPorId("1")

	if erro != nil {
		t.Errorf("Erro ao buscar administrasdor por id %s", erro.Error())
	}

	esperado := reflect.TypeOf(models.Administrador{})

	if reflect.TypeOf(*admDb) != esperado {
		t.Errorf("Tipo esperado: %v, obtido: %v", esperado, reflect.TypeOf(admDb))
	}

	if admDb.Id != "" {
		t.Errorf("O administrador não deveria existir")
	}
}

func TestApagar(t *testing.T) {
	db, _ := config.GetDB()
	truncateTable(db)

	repo := GenericoRepositorioMySql[models.Administrador]{
		DB: db,
	}

	adm := models.Administrador{
		Id:    "1",
		Nome:  "teste",
		Email: "teste",
		Senha: "teste",
	}

	erroAdd := repo.Adicionar(adm)
	if erroAdd != nil {
		t.Errorf("Erro ao cadastrar administrador %s", erroAdd.Error())
	}

	erroApagar := repo.ApagarPorId("1")
	if erroApagar != nil {
		t.Errorf("Erro ao apagar administrador %s", erroApagar.Error())
	}

	adms, erro := repo.Lista()

	if erro != nil {
		t.Errorf("Erro ao buscar administrasdor por id %s", erro.Error())
	}

	if len(adms) != 0 {
		t.Errorf("Ao apagar s tabela de administradores deveria vir 0")
	}
}

func TestAtualizar(t *testing.T) {
	db, _ := config.GetDB()
	truncateTable(db)

	repo := GenericoRepositorioMySql[models.Administrador]{
		DB: db,
	}

	adm := models.Administrador{
		Id:    "1",
		Nome:  "teste",
		Email: "teste",
		Senha: "teste",
	}

	erroAdd := repo.Adicionar(adm)
	if erroAdd != nil {
		t.Errorf("Erro ao cadastrar administrador %s", erroAdd.Error())
	}

	admDb, _ := repo.BuscaPorId("1")

	admDb.Nome = "danilo"
	admDb.Email = "danilo@teste.com"
	admDb.Senha = "123456"
	admDb.Super = true

	erroAlterar := repo.Alterar(*admDb)
	if erroAlterar != nil {
		t.Errorf("Erro ao atualizar administrasdor %s", erroAlterar.Error())
	}

	admDbAtualizado, _ := repo.BuscaPorId("1")

	if admDbAtualizado.Nome != admDb.Nome {
		t.Errorf("O Nome não foi atualizado")
	}

	if admDbAtualizado.Email != admDb.Email {
		t.Errorf("O Email não foi atualizado")
	}

	if admDbAtualizado.Senha != admDb.Senha {
		t.Errorf("O Senha não foi atualizado")
	}

	if admDbAtualizado.Super != admDb.Super {
		t.Errorf("O Super não foi atualizado")
	}

}
