package services

import (
	"http_fiber_api/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Definição do claim do token
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken gera um token JWT para um administrador
func GenerateToken(email string) (string, error) {
	// Definir as claims do token, incluindo o email do administrador e a validade do token
	claims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 horas de validade
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Criar o token com as claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assinar o token com a chave secreta
	tokenString, err := token.SignedString(config.GetJWTSecretKey())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
