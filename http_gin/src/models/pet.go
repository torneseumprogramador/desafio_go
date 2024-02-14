package models

import "http_gin/src/enums"

type Pet struct {
	Id   string
	Nome string
	Dono string
	Tipo enums.Tipo
}
