package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

// func TestMain(m *testing.M) {
// 	// Inicia o servidor em uma goroutine
// 	go func() {
// 		log.Println("Iniciando servidor para testes...")
// 		startServer()
// 	}()

// 	// Dá um tempo para o servidor iniciar
// 	log.Println("Aguardando o servidor iniciar...")
// 	time.Sleep(2 * time.Second)

// 	// Executa os testes
// 	code := m.Run()

// 	// Aqui você pode adicionar código para desligar o servidor se necessário
// 	// Por exemplo, se você tiver uma função stopServer(), chame-a aqui

// 	// Sair com o código de status dos testes
// 	os.Exit(code)
// }

func TestHelloEndpointIntegration(t *testing.T) {
	// Faz a requisição
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("Não foi possível fazer a requisição GET: %v", err)
	}
	defer resp.Body.Close()

	// Checa o status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code errado, recebido: %d, esperado: %d", resp.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Erro ao ler o corpo da resposta: %v", err)
	}
	bodyString := string(body)

	// Checa o corpo da resposta
	expected := "Teste de request no home\n" // Note o '\n' no final, que é adicionado por fmt.Fprintln
	if bodyString != expected {
		t.Errorf("Corpo errado, recebido: %s, esperado: %s", bodyString, expected)
	}
}

func TestPostRequest(t *testing.T) {
	// Cria o corpo da requisição
	pet := Pet{
		Id:   "1",
		Nome: "Rex",
		Dono: "João",
	}
	body, err := json.Marshal(pet)
	if err != nil {
		t.Fatalf("Não foi possível serializar o pet: %v", err)
	}

	// Faz a requisição POST
	resp, err := http.Post("http://localhost:8080/pets", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Não foi possível fazer a requisição POST: %v", err)
	}
	defer resp.Body.Close()

	// Checa o status code
	if resp.StatusCode != http.StatusCreated { // Use http.StatusCreated para corresponder ao status esperado
		t.Errorf("Status code errado, recebido: %d, esperado: %d", resp.StatusCode, http.StatusCreated)
	}

	// Lê o corpo da resposta
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Erro ao ler o corpo da resposta: %v", err)
	}
	bodyString := string(respBody)

	// Checa o corpo da resposta
	expected := `{"mensagem": "Pet 'Rex' recebido com sucesso"}` // Ajusta conforme a resposta esperada do handler
	if bodyString != expected {
		t.Errorf("Corpo errado, recebido: %s, esperado: %s", bodyString, expected)
	}
}
