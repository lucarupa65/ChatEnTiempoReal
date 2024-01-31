// main.go
package main

import (
	"log"

	"github.com/luisk6510/ChatEnTiempoReal/db"
	"github.com/luisk6510/ChatEnTiempoReal/internal/app"
	"github.com/luisk6510/ChatEnTiempoReal/web/di"
)

func main() {
	// Configurar la conexión a MongoDB
	cadenaConexion := "mongodb://localhost:27017"
	nombreBaseDatos := "miappdb"
	mongoDB, err := db.ObtenerInstanciaMongoDB(cadenaConexion, nombreBaseDatos)
	if err != nil {
		log.Fatal("Error al obtener la instancia de MongoDB:", err)
	}
	defer mongoDB.CerrarConexion()

	// Construir el contenedor de dependencias
	container := di.BuildContainer()

	// Resolver las dependencias
	err = container.Invoke(func(app *app.Application) {
		// Lanzar la aplicación
		app.Run(mongoDB)
	})
	if err != nil {
		log.Fatal("Error al construir y resolver dependencias:", err)
	}
}
