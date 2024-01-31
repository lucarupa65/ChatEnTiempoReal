// user.go
package domain

// User representa la entidad de usuario en el dominio del negocio.
type User struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	// Otros campos del usuario...
}
