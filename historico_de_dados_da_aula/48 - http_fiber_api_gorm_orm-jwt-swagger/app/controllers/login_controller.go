package controllers

import (
	"http_fiber_api/app/dtos"
	modelviews "http_fiber_api/app/model_views"
	"http_fiber_api/app/models"
	"http_fiber_api/app/services"
	"http_fiber_api/plataform/database"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct{}

// Login godoc
// @Summary Login do Administrador
// @Description Autentica um administrador e retorna um token JWT
// @Tags Login
// @Accept json
// @Produce json
// @Param loginDto body dtos.LoginDto true "Informações de Login"
// @Success 200 {object} modelviews.AdmToken "Token de autenticação"
// @Failure 400 {object} map[string]string "Erro ao analisar o corpo da requisição"
// @Failure 401 {object} map[string]string "Credenciais inválidas"
// @Router /login [post]
func (pc *LoginController) Login(c *fiber.Ctx) error {
	loginDto := new(dtos.LoginDto)

	if err := c.BodyParser(loginDto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"mensagem": "Erro ao analisar o corpo da requisição - " + err.Error(),
		})
	}

	var adm models.Administrador
	db := database.GetConnection()
	result := db.Where("email = ? AND senha = ?", loginDto.Email, loginDto.Senha).First(&adm)

	if result.Error != nil {
		return c.Status(401).JSON(fiber.Map{
			"mensagem": "Credenciais inválidas",
		})
	}

	token, erroToken := services.GenerateToken(adm.Email)

	if erroToken != nil {
		return c.Status(400).JSON(fiber.Map{
			"mensagem": erroToken.Error(),
		})
	}

	admToken := modelviews.AdmToken{
		Id:    adm.ID,
		Nome:  adm.Nome,
		Email: adm.Email,
		Token: token,
	}

	return c.Status(200).JSON(admToken)
}
