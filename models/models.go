package models

import "time"

type Nota struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Titulo    string    `json:"titulo"`
	Conteudo  string    `json:"conteudo"`
	CreatedAt time.Time `json:"criado_em"`
	UpdatedAt time.Time `json:"atualizado_em"`
}
