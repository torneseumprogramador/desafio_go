package models

import (
	"time"
)

type Administrador struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Nome      string     `gorm:"type:varchar(150);not null" json:"nome"`
	Email     string     `gorm:"type:varchar(255);not null" json:"email"`
	Senha     string     `gorm:"type:varchar(100);not null" json:"senha"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

func (Administrador) TableName() string {
	return "administradores"
}
