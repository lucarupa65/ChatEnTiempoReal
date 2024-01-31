// application.go
package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luisk6510/ChatEnTiempoReal/db"
	"github.com/luisk6510/ChatEnTiempoReal/internal/repository"
	"github.com/luisk6510/ChatEnTiempoReal/internal/service"
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
func (app *Application) Run(mongoDB *db.MongoDB) {
	// Inicializar y configurar servicios, repositorios, etc.
	// ...

	// Iniciar el servidor HTTP
	app.iniciarServidor(mongoDB)
}

// iniciarServidor inicia el servidor HTTP.
func (app *Application) iniciarServidor(mongoDB *db.MongoDB) {
	// Inicializar y configurar servicios, repositorios, etc.
	userRepository := repository.NewUserRepository(mongoDB)
	userService := service.NewUserService(*userRepository)

	// Crear el controlador de usuario
	userHandler := handler.NewUserHandler(userService)

	// Configurar rutas HTTP
	http.HandleFunc("/usuarios", userHandler.CrearUsuarioHandler)

	// Iniciar el servidor
	puerto := 8080
	direccion := fmt.Sprintf(":%d", puerto)
	fmt.Printf("Servidor iniciado en http://localhost:%d\n", puerto)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
