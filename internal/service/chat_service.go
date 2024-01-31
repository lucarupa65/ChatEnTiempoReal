package service

import "github.com/luisk6510/ChatEnTiempoReal/internal/repository"

// UserService representa un servicio de usuario.
type ChatService struct {
	ChatRepo repository.ChatRepository
}

// NewUserService crea una nueva instancia de UserService.
func NewChatService(chatRepo repository.ChatRepository) *ChatService {
	return &ChatService{
		ChatRepo: chatRepo,
	}
}
