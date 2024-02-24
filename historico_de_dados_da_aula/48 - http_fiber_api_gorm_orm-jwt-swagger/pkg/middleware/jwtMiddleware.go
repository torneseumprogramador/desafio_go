package middleware

import (
	"http_fiber_api/pkg/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: config.GetJWTSecretKey(),
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Não autorizado",
			})
		},
	})
}
