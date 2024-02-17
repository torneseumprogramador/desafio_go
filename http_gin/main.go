package main

import (
	"http_gin/src/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Carrega os templates HTML
	r.LoadHTMLGlob("src/templates/**/*.tmpl.html")

	r.Static("/public/", "./public_assets")

	config.Routes(r)

	r.Run(":5000") // Por padrão, escuta na porta 5000
}

// package main

// import (
// 	"fmt"
// 	"http_gin/src/database"
// 	"http_gin/src/models"
// 	"http_gin/src/repositorios"
// 	"http_gin/src/servicos"
// 	"log"
// 	"reflect"
// )

// // https://chat.openai.com/share/5b0138b6-e0d5-4fe5-98cb-fd87c9cbcb8a
// // https://chat.openai.com/share/6ed8365b-6b91-48a1-a5ad-7d4b3ae214d5
// func main() {
// 	db, err := database.GetDB()
// 	if err != nil {
// 		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
// 	}
// 	defer db.Close()

// 	administradorRepo := repositorios.AdministradorRepositorioMySql{DB: db}
// 	servico := servicos.NovoCrudServico[models.Administrador](&administradorRepo)

// 	// Exemplo de uso do serviço
// 	administradores, err := servico.Repo.Lista()
// 	if err != nil {
// 		log.Fatalf("Erro ao listar administradores: %v", err)
// 	}
// 	log.Println("Administradores encontrados:", administradores)

// 	fmt.Println("================ [Generico Adm] ====================")

// 	adminGenericoRepo := repositorios.GenericoRepositorioMySql{
// 		DB:    db,
// 		Table: "administradores",
// 		Type:  reflect.TypeOf(models.Administrador{}),
// 	}

// 	// repoGenerico := servicos.NovoRepositorio[models.Administrador](&adminGenericoRepo)
// 	// adms, err := repoGenerico.Servico.Lista()

// 	adms, err := adminGenericoRepo.Lista()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Administradores encontrados:", adms)

// 	fmt.Println("================ [Generico Dono] ====================")

// 	donoGenericoRepo := repositorios.GenericoRepositorioMySql{
// 		DB:    db,
// 		Table: "donos",
// 		Type:  reflect.TypeOf(models.Dono{}),
// 	}

// 	donos, err := donoGenericoRepo.Lista()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Donos encontrados:", donos)

// }
