package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Pet struct {
	Id   string `json:"id"`
	Nome string `json:"nome"`
	Dono string `json:"dono"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Teste de request no home")
}

func petsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var pet Pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Aqui você pode processar o pet como necessário, por exemplo, salvá-lo em um banco de dados.

	// Envia uma resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created é o código de status apropriado para uma criação bem sucedida.
	fmt.Fprintf(w, `{"mensagem": "Pet '%s' recebido com sucesso"}`, pet.Nome)
}

func startServer() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/pets", petsHandler)
	fmt.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	startServer()
}
