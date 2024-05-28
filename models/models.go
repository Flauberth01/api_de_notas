package models

import "time"

type Nota struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Titulo       string    `json:"titulo"`
	Conteudo     string    `json:"conteudo"`
	CriadoEm     time.Time `json:"criado_em"`
	AtualizadoEm time.Time `json:"atualizado_em"`
}
