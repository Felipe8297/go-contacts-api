package contacts

import (
	"time"
)

// @Description Informações de um contato
type Contact struct {
	ID         string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`          // ID único do contato
	Name       string    `json:"name" example:"João Silva"`                                  // Nome do contato
	Email      string    `json:"email" example:"joao@example.com"`                           // Email do contato
	Phone      string    `json:"phone" example:"11999998888"`                                // Telefone do contato
	CategoryID string    `json:"category_id" example:"123e4567-e89b-12d3-a456-426614174111"` // ID da categoria
	CreatedAt  time.Time `json:"created_at" example:"2023-01-01T12:00:00Z"`                  // Data de criação
	UpdatedAt  time.Time `json:"updated_at" example:"2023-01-01T12:00:00Z"`                  // Data de atualização
}
