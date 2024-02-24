package controllers

import (
	modelviews "http_fiber_api/app/model_views"

	"github.com/gofiber/fiber/v2"
)

type HomeController struct{}

// Index godoc
// @Summary Apresentação da API
// @Description Mostrar dados sobre a API quando acessar a rota /
// @Tags Home
// @Accept  json
// @Produce  json
// @Success 200 {object} modelviews.HomeView
// @Router / [get]
func (pc *HomeController) Index(c *fiber.Ctx) error {
	return c.Status(200).JSON(modelviews.HomeView{
		Mensagem: "API com Fiber e GORM feito no desafio de Golang",
		Doc:      "/swagger/index.html",
	})
}
