package repository

import "github.com/luisk6510/ChatEnTiempoReal/db"

// ChatRepository representa un repositorio de los chat.
type ChatRepository struct {
	DB *db.MongoDB
}

// NewChatRepository crea una nueva instancia de ChatRepository.
func NewChatRepository(db *db.MongoDB) *ChatRepository {
	return &ChatRepository{
		DB: db,
	}
}
