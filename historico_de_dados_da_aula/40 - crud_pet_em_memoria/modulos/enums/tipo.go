package enums

type Tipo int

// Definindo os valores do enum
const (
	Cachorro Tipo = iota // iota facilita a atribuição incremental de valores
	Gato
	Outros
)
