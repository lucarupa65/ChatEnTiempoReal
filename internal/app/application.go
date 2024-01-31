// application.go
package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luisk6510/ChatEnTiempoReal/config"
	"github.com/luisk6510/ChatEnTiempoReal/web/handler"
)

// Application representa la aplicación principal.
type Application struct {
	UserHandler *handler.UserHandler
}

// NewApplication crea una nueva instancia de Application.
func NewApplication(userHandler *handler.UserHandler) *Application {
	return &Application{
		UserHandler: userHandler,
	}
}

// Run inicia la aplicación.
func (app *Application) Run(serverConfig config.ServerConfig) {
	// Inicializar y configurar servicios, repositorios, etc.
	// ...

	// Iniciar el servidor HTTP
	app.iniciarServidor(serverConfig)
}

// iniciarServidor inicia el servidor HTTP.
func (app *Application) iniciarServidor(serverConfig config.ServerConfig) {

	// Configurar rutas HTTP
	http.HandleFunc("/usuarios", app.UserHandler.CrearUsuarioHandler)

	// Iniciar el servidor
	direccion := fmt.Sprintf(":%d", serverConfig.Port)
	fmt.Printf("Servidor iniciado en http://localhost:%d\n", serverConfig.Port)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
