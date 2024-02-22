package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HomeController struct{}

func (pc *HomeController) Index(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"mensagem": "API com fiber",
	})
}
