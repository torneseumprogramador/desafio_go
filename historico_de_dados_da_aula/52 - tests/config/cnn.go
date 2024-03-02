// config/cnn.go

package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GetDB retorna uma conexão com o banco de dados
func GetDB() (*sql.DB, error) {
	fmt.Println("=====================")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "root")
	dbHost := getEnv("DB_HOST", "127.0.0.1:3306")
	dbName := getEnv("DB_NAME", "desafio_go")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Erro ao abrir a conexão com o banco de dados: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	return db, nil
}
