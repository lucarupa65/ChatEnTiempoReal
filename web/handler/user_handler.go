// user_handler.go
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/luisk6510/ChatEnTiempoReal/internal/service"
)

// UserHandler maneja las solicitudes relacionadas con los usuarios.
type UserHandler struct {
	UserService *service.UserService
}

// NewUserHandler crea una nueva instancia de UserHandler.
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// CrearUsuarioHandler maneja la creación de un nuevo usuario.
func (h *UserHandler) CrearUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	// Parsear los datos del cuerpo de la solicitud
	var nuevoUsuario struct {
		Nombre string `json:"nombre"`
		// Otros campos del usuario...
	}

	if err := json.NewDecoder(r.Body).Decode(&nuevoUsuario); err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Crear el usuario utilizando el servicio
	err := h.UserService.CrearUsuario(r.Context(), nuevoUsuario.Nombre)
	if err != nil {
		http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusCreated)
}
