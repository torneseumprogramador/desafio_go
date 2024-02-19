package models

//lint:ignore U1000 reason: Used by ORM to specify table name
type Administrador struct {
	tableName struct{} `table:"administradores"`

	Id    string `field:"id"`
	Nome  string `field:"nome"`
	Email string `field:"email"`
	Senha string `field:"senha"`
}
