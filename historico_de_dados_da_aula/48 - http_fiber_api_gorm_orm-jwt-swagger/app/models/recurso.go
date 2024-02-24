package models

import (
	"time"
)

type Recurso struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Nome      string     `gorm:"type:varchar(100);not null" json:"nome"`
	Valor     float32    `gorm:"type:decimal(10,2)" json:"valor"`
}

func (Recurso) TableName() string {
	return "recursos"
}
