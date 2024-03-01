package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" // Importe o pacote go-redis/redis
)

type Pagina struct {
	Id       int
	Titulo   string
	Conteudo string
}

func buscandoListaDoBancoDeDados() []Pagina {
	paginas := []Pagina{
		Pagina{Id: 1, Titulo: "Home", Conteudo: "<html><head><title>Home page</title><body> ..."},
		Pagina{Id: 2, Titulo: "Login", Conteudo: "<html><head><title>Login</title><body> ..."},
		Pagina{Id: 3, Titulo: "Dashboard", Conteudo: "<html><head><title>Dashboard</title><body> ..."},
	}

	return paginas
}

func jsonParse(paginas []Pagina) (string, error) {
	bytes, err := json.Marshal(paginas)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func main() {
	// Crie um novo cliente Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // endereço do servidor Redis
		Password: "",               // senha, se houver
		DB:       0,                // banco de dados padrão a ser usado
	})

	ctx := context.Background()

	var paginas []Pagina

	listaPaginaJson, _ := rdb.Get(ctx, "lista-paginas").Result()
	if listaPaginaJson == "" {
		paginas = buscandoListaDoBancoDeDados()
		stringJson, _ := jsonParse(paginas)

		fmt.Println("=======Gravando no redis===========")
		err := rdb.Set(ctx, "lista-paginas", stringJson, 10*time.Second).Err() // com 10 segundos de cache
		if err != nil {
			panic(err)
		}
		fmt.Println("=======gravação ok===========")
	} else {
		err := json.Unmarshal([]byte(listaPaginaJson), &paginas)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	for _, pagina := range paginas {
		fmt.Printf("ID: %d\nTítulo: %s\nConteúdo: %s\n", pagina.Id, pagina.Titulo, pagina.Conteudo)
		fmt.Println("--------------------------------------------------")
	}

	// Definindo uma chave "key" com o valor "value" no Redis
	// fmt.Println("=======Gravando no redis===========")
	// err := rdb.Set(ctx, "key", "tester 1234567", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("=======gravação ok===========")

	// // Obtendo o valor da chave "key" do Redis
	// fmt.Println("=======Buscando do redis===========")
	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Item no redis=", val)

	// // Tente obter o valor de uma chave que não existe
	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
}
