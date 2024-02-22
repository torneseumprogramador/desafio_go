package models

import (
	"gorm.io/gorm"
)

type Recurso struct {
	gorm.Model
	Nome  string  `gorm:"type:varchar(100);not null" json:"nome"`
	Valor float32 `gorm:"type:decimal(10,2)" json:"valor"`
}

func (Recurso) TableName() string {
	return "recursos"
}
