// user_service.go
package service

import (
	"context"
	"github.com/luisk6510/ChatEnTiempoReal/internal/domain"
	"github.com/luisk6510/ChatEnTiempoReal/internal/repository"
)

// UserService representa un servicio de usuario.
type UserService struct {
	UserRepo repository.UserRepository
}

// NewUserService crea una nueva instancia de UserService.
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

// CrearUsuario crea un nuevo usuario.
func (s *UserService) CrearUsuario(ctx context.Context, nombre string) error {
	usuario := domain.User{
		Nombre: nombre,
		// Otros campos del usuario...
	}

	return s.UserRepo.GuardarUsuario(ctx, usuario)
}
