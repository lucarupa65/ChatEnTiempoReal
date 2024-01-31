// main.go
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/luisk6510/ChatEnTiempoReal/config"
	"github.com/luisk6510/ChatEnTiempoReal/internal/app"
	"github.com/luisk6510/ChatEnTiempoReal/web/di"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	//Obtenemos las variables de configuracion
	configu := config.NewConfig()

	// Construir el contenedor de dependencias
	container := di.BuildContainer()

	// Resolver las dependencias
	err = container.Invoke(func(app *app.Application) {
		// Lanzar la aplicación
		app.Run(configu.Serve)
	})
	if err != nil {
		log.Fatal("Error al construir y resolver dependencias:", err)
	}
}
