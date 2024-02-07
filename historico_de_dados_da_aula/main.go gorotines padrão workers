package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker que processa tarefas
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Notifica a conclusão ao WaitGroup
	for j := range jobs {
		fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simula um trabalho sendo feito
		fmt.Printf("worker %d finished job %d\n", id, j)
		results <- j * 2 // Envia o resultado do "processamento"
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Inicia um número fixo de workers.
	const numWorkers = 3
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1) // Adiciona um contador para cada worker
		go worker(w, jobs, results, &wg)
	}

	// Envia tarefas para o canal de jobs.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Fecha o canal de jobs para indicar que não há mais tarefas.

	// Espera todos os workers terminarem.
	wg.Wait()

	// Coleta os resultados das tarefas processadas.
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	close(results) // Opcional, fecha o canal de results
	fmt.Println("Todas as tarefas foram processadas.")
}
