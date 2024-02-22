package models

type Recurso struct {
	Id    uint32  `db:"id"`
	Nome  string  `db:"nome"`
	Valor float32 `db:"valor"`
}
