package middleware

import (
	"http_fiber_api/pkg/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ProtectedJwtLowLevel() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Cabeçalho Authorization não encontrado",
			})
		}

		// Dividir o cabeçalho Authorization para obter o token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Formato do cabeçalho Authorization inválido",
			})
		}

		// Aqui você tem o token
		tokenString := parts[1]

		// Decodificar e validar o token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Certifique-se de que o algoritmo de assinatura esperado seja utilizado!
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}

			return config.GetJWTSecretKey(), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido",
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido ou expirado",
			})
		}

		// Token é válido; prosseguir com a requisição
		return c.Next()
	}
}
