package controllers

import (
	"fmt"
	"http_gin/src/DTOs"
	"http_gin/src/libs"
	"http_gin/src/models"
	"http_gin/src/servicos"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type LoginController struct{}

func (hc *LoginController) Login(c *gin.Context) {
	var loginDTO DTOs.LoginDTO

	if err := c.BindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())

	credenciais := make(map[string]string)
	credenciais["email"] = loginDTO.Email

	adms, erro := servico.Repo.Where(credenciais)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	if len(adms) > 0 && libs.CryptoEq(loginDTO.Senha, adms[0].Senha) {
		adm := adms[0]
		token, erro := tokenJwt(c, adm)
		if erro != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": erro.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"id":    adm.Id,
			"nome":  adm.Nome,
			"email": adm.Email,
			"super": adm.Super,
		})

		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Email ou senha inválido",
	})
}

func tokenJwt(c *gin.Context, adm models.Administrador) (string, error) {
	tempoExpiracao := time.Now().Add(time.Hour * 1)

	token := jwt.New(jwt.SigningMethodHS256)

	// Define claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = adm.Id
	claims["nome"] = adm.Nome
	claims["email"] = adm.Email
	claims["super"] = adm.Super
	claims["exp"] = tempoExpiracao.Unix()

	chave := libs.GetEnv("JWT_TOKEN", "desafio_go")
	tokenString, err := token.SignedString([]byte(chave))
	if err != nil {
		return "", fmt.Errorf("Login ou senha inválido")
	}

	return tokenString, nil
}
