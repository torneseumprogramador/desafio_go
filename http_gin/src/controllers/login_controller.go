package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"http_gin/src/libs"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"http_gin/src/servicos"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type LoginController struct{}

func (hc *LoginController) Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Página de login",
			"currentRoute": "login",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/Login/index.tmpl.html",
					nil,
				),
			),
		},
	)
}

func (hc *LoginController) Sair(c *gin.Context) {
	c.SetCookie("adminInfo", "", -1, "/", "", false, true)
	c.Redirect(302, "/login")
}

func (hc *LoginController) Login(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())

	credenciais := make(map[string]string)
	credenciais["email"] = c.Request.FormValue("email")

	adms, erro := servico.Repo.Where(credenciais)

	if len(adms) > 0 && libs.CryptoEq(c.Request.FormValue("senha"), adms[0].Senha) {
		adm := adms[0]

		admView := model_views.AdmView{
			Id:    adm.Id,
			Email: adm.Email,
			Nome:  adm.Nome,
		}

		cookieValueBytes, err := json.Marshal(admView)
		if err != nil {
			c.Redirect(302, "/login")
			return
		}

		cookieValue := string(cookieValueBytes)
		encodedValue := url.QueryEscape(cookieValue)

		tempoExpiracao := time.Now().Add(time.Hour * 1)

		token := jwt.New(jwt.SigningMethodHS256)

		// Define claims
		claims := token.Claims.(jwt.MapClaims)
		claims["sub"] = encodedValue
		claims["exp"] = tempoExpiracao.Unix()

		chave := libs.GetEnv("JWT_TOKEN", "desafio_go")
		tokenString, err := token.SignedString([]byte(chave))
		if err != nil {
			c.Redirect(302, "/login")
			return
		}

		duracao := int(tempoExpiracao.Unix() - time.Now().Unix())

		c.SetCookie("adminInfo", tokenString, duracao, "/", "", false, true)

		c.Redirect(302, "/pets")
		return
	}

	if erro == nil {
		erro = fmt.Errorf("Login ou senha inválido")
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Página de login",
			"currentRoute": "login",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/Login/index.tmpl.html",
					map[string]interface{}{
						"erro": erro,
					},
				),
			),
		},
	)
}
