package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	// Inicia o servidor em uma goroutine
	go func() {
		log.Println("Iniciando servidor para testes...")
		startServer()
	}()

	// Dá um tempo para o servidor iniciar
	log.Println("Aguardando o servidor iniciar...")
	time.Sleep(2 * time.Second)

	// Executa os testes
	code := m.Run()

	// Aqui você pode adicionar código para desligar o servidor se necessário
	// Por exemplo, se você tiver uma função stopServer(), chame-a aqui

	// Sair com o código de status dos testes
	os.Exit(code)
}

func TestAPIBehavior(t *testing.T) {
	Convey("Dado um servidor HTTP rodando", t, func() {
		Convey("Quando o endpoint GET / é chamado", func() {
			resp, err := http.Get("http://localhost:8080")
			So(err, ShouldBeNil)
			defer resp.Body.Close()

			Convey("Então o status code deve ser 200 OK", func() {
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
			})

			Convey("E o corpo da resposta deve ser 'Teste de request no home\n'", func() {
				body, _ := io.ReadAll(resp.Body)
				So(string(body), ShouldEqual, "Teste de request no home\n")
			})
		})

		Convey("Quando o endpoint POST /pets é chamado com um novo pet", func() {
			pet := Pet{Id: "1", Nome: "Rex", Dono: "João"}
			body, _ := json.Marshal(pet)
			resp, err := http.Post("http://localhost:8080/pets", "application/json", bytes.NewReader(body))
			So(err, ShouldBeNil)
			defer resp.Body.Close()

			Convey("Então o status code deve ser 201 Created", func() {
				So(resp.StatusCode, ShouldEqual, http.StatusCreated)
			})

			Convey("E a resposta deve confirmar o recebimento do pet", func() {
				respBody, _ := io.ReadAll(resp.Body)
				So(string(respBody), ShouldEqual, `{"mensagem": "Pet 'Rex' recebido com sucesso"}`)
			})
		})
	})
}
