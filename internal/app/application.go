// application.go
package app

import (
	"fmt"
	"log"
	"net/http"

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
func (app *Application) Run() {
	// Inicializar y configurar servicios, repositorios, etc.
	// ...

	// Iniciar el servidor HTTP
	app.iniciarServidor()
}

// iniciarServidor inicia el servidor HTTP.
func (app *Application) iniciarServidor() {

	// Configurar rutas HTTP
	http.HandleFunc("/usuarios", app.UserHandler.CrearUsuarioHandler)

	// Iniciar el servidor
	puerto := 8080
	direccion := fmt.Sprintf(":%d", puerto)
	fmt.Printf("Servidor iniciado en http://localhost:%d\n", puerto)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
