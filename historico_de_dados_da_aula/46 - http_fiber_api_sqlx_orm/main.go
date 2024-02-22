// package main

// import (
// 	"http_fiber_api/pkg/routes"

// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()
// 	routes.RegisterPublicRoutes(app)
// 	app.Listen(":3000")
// }

package main

import (
	"fmt"
	"http_fiber_api/app/models"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func execMigrations(db *sqlx.DB) error {
	// Definir o diretório das migrations
	migrationsDir := "./plataform/migrations"

	// Listar todos os arquivos SQL no diretório
	err := filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".sql" {
			// Ler o conteúdo do arquivo SQL
			migration, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Executar a migration
			if _, err = db.Exec(string(migration)); err != nil {
				return fmt.Errorf("error executing migration %s: %w", path, err)
			}
			fmt.Printf("Migration executed successfully: %s\n", path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return nil
}

func ConnectToDB() (*sqlx.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/http_fiber_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateRecurso(db *sqlx.DB, r *models.Recurso) error {
	query := `INSERT INTO recursos (nome, valor) VALUES (:nome, :valor)`
	_, err := db.NamedExec(query, r)
	return err
}

func GetAllRecursos(db *sqlx.DB) ([]models.Recurso, error) {
	var recursos []models.Recurso
	query := `SELECT id, nome, valor FROM recursos`
	err := db.Select(&recursos, query)
	return recursos, err
}

func GetRecursoById(db *sqlx.DB, id uint32) (*models.Recurso, error) {
	var recurso models.Recurso
	query := `SELECT id, nome, valor FROM recursos WHERE id = ?`
	err := db.Get(&recurso, query, id)
	if err != nil {
		return nil, err
	}
	return &recurso, nil
}

func UpdateRecurso(db *sqlx.DB, r *models.Recurso) error {
	query := `UPDATE recursos SET nome = :nome, valor = :valor WHERE id = :id`
	_, err := db.NamedExec(query, r)
	return err
}

func DeleteRecurso(db *sqlx.DB, id uint32) error {
	query := `DELETE FROM recursos WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

func main() {
	db, err := ConnectToDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer db.Close()

	fmt.Println("==== Conexão realizada com sucesso =====")
	fmt.Println(db)

	execMigrations(db)

	println("Migration executada com sucesso!")
	fmt.Println("==== Migration criada com sucesso =====")

	// Criar um novo Recurso
	newRecurso := models.Recurso{
		Nome:  "Exemplo Recurso",
		Valor: 100.50,
	}
	if err := CreateRecurso(db, &newRecurso); err != nil {
		log.Fatalf("Erro ao criar recurso: %v", err)
	}
	fmt.Println("Recurso criado com sucesso!")

	// Buscar todos os Recursos
	recursos, err := GetAllRecursos(db)
	if err != nil {
		log.Fatalf("Erro ao buscar recursos: %v", err)
	}
	fmt.Println("Recursos encontrados:", recursos)

	// Buscar um Recurso por ID
	// Substitua 1 pelo ID do recurso que você deseja buscar
	recurso, err := GetRecursoById(db, 1)
	if err != nil {
		log.Fatalf("Erro ao buscar recurso por ID: %v", err)
	}
	fmt.Printf("Recurso encontrado: %+v\n", recurso)

	// Atualizar um Recurso
	// Substitua 1 pelo ID do recurso que você deseja atualizar
	recursoParaAtualizar := models.Recurso{
		Id:    1, // Assumindo que o recurso com ID 1 existe
		Nome:  "Recurso Atualizado",
		Valor: 200.75,
	}
	if err := UpdateRecurso(db, &recursoParaAtualizar); err != nil {
		log.Fatalf("Erro ao atualizar recurso: %v", err)
	}
	fmt.Println("Recurso atualizado com sucesso!")

	// Excluir um Recurso
	// Substitua 1 pelo ID do recurso que você deseja excluir
	if err := DeleteRecurso(db, 1); err != nil {
		log.Fatalf("Erro ao excluir recurso: %v", err)
	}
	fmt.Println("Recurso excluído com sucesso!")

	///===== Exercicio para vc aluno, usar os métodos de CRUD acima colocando no controlador de recursos_controller.go ========
}
