// main.go
package main

import (
	"log"

	"github.com/luisk6510/ChatEnTiempoReal/internal/app"
	"github.com/luisk6510/ChatEnTiempoReal/web/di"
)

func main() {

	// Construir el contenedor de dependencias
	container := di.BuildContainer()

	// Resolver las dependencias
	err := container.Invoke(func(app *app.Application) {
		// Lanzar la aplicación
		app.Run()
	})
	if err != nil {
		log.Fatal("Error al construir y resolver dependencias:", err)
	}
}
