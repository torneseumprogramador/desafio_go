package aniversario

import "time"

func QuantosDiasFaltamParaSeuAniversario(nome string, dataNascimento string) (string, int, int, int) {
	// Converte a string de data para o tipo time.Time
	layout := "02/01/2006" // Layout usado para parsear a data, Go usa esse formato específico
	nascimento, _ := time.Parse(layout, dataNascimento)

	// Calcula a idade
	hoje := time.Now()
	idade := hoje.Year() - nascimento.Year()
	if hoje.YearDay() < nascimento.YearDay() { // Ainda não fez aniversário este ano
		idade--
	}

	// Calcula quantos dias e horas faltam para o próximo aniversário
	proximoAniversario := time.Date(hoje.Year(), nascimento.Month(), nascimento.Day(), 0, 0, 0, 0, time.Local)
	if proximoAniversario.Before(hoje) {
		proximoAniversario = proximoAniversario.AddDate(1, 0, 0) // Próximo ano
	}
	diferenca := proximoAniversario.Sub(hoje)
	diasFaltamAniversario := int(diferenca.Hours() / 24)
	horasFaltamAniversario := int(diferenca.Hours()) % 24 // Calcula as horas restantes após os dias inteiros

	return "Oi - " + nome, idade, diasFaltamAniversario, horasFaltamAniversario
}
