package domain

// ChatRoom representa la entidad de una sala de chat en el dominio del negocio.
type ChatRoom struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	// Otros campos de la sala de chat...
}

// Mensaje representa la entidad de un mensaje en el dominio del negocio.
type Mensaje struct {
	ID        string `json:"id"`
	Contenido string `json:"contenido"`
	Remitente User   `json:"remitente"`
	// Otros campos del mensaje...
}
