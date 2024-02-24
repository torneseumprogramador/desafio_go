package config

import "os"

func GetJWTSecretKey() []byte {
	// Tente obter o valor da variável de ambiente
	secret := os.Getenv("JWT_SECRET_KEY")

	// Se não estiver definido, use o valor padrão
	if secret == "" {
		secret = "desafio_go_token_jwt"
	}

	return []byte(secret)
}

func GetDSN() string {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = "root:root@tcp(localhost:3306)/http_fiber_api?charset=utf8mb4&parseTime=True&loc=Local"
	}

	return dsn
}
