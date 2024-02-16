// package main

// import (
// 	"http_gin/src/config"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	// Carrega os templates HTML
// 	r.LoadHTMLGlob("src/templates/**/*.tmpl.html")

// 	r.Static("/public/", "./public_assets")

// 	config.Routes(r)

// 	r.Run(":5000") // Por padrão, escuta na porta 5000
// }

package main

import (
	"fmt"
	"http_gin/src/database"
	"http_gin/src/models"
	"http_gin/src/servicos"
	"log"
	"reflect"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	administradorServico := servicos.AdministradorServicoMySql{DB: db}
	repo := servicos.NovoRepositorio[models.Administrador](&administradorServico)

	// Exemplo de uso do serviço
	administradores, err := repo.Servico.Lista()
	if err != nil {
		log.Fatalf("Erro ao listar administradores: %v", err)
	}
	log.Println("Administradores encontrados:", administradores)

	fmt.Println("================ [Generico Adm] ====================")

	adminGenericoService := servicos.GenericService{
		DB:    db,
		Table: "administradores",
		Type:  reflect.TypeOf(models.Administrador{}),
	}

	// repoGenerico := servicos.NovoRepositorio[models.Administrador](&adminGenericoService)
	// adms, err := repoGenerico.Servico.Lista()

	adms, err := adminGenericoService.Lista()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Administradores encontrados:", adms)

	fmt.Println("================ [Generico Dono] ====================")

	donoService := servicos.GenericService{
		DB:    db,
		Table: "donos",
		Type:  reflect.TypeOf(models.Dono{}),
	}

	donos, err := donoService.Lista()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Donos encontrados:", donos)

}
