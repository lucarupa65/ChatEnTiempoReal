// container.go
package di

import (
	"github.com/luisk6510/ChatEnTiempoReal/db"
	"github.com/luisk6510/ChatEnTiempoReal/internal/app"
	"github.com/luisk6510/ChatEnTiempoReal/internal/repository"
	"github.com/luisk6510/ChatEnTiempoReal/internal/service"
	"github.com/luisk6510/ChatEnTiempoReal/web/handler"

	"go.uber.org/dig"
)

var Container *dig.Container

// BuildContainer construye el contenedor de dependencias.
func BuildContainer() *dig.Container {
	Container = dig.New()

	// Configurar la conexión a MongoDB
	cadenaConexion := "mongodb://localhost:27017"
	nombreBaseDatos := "miappdb"
	mongoDB, err := db.ObtenerInstanciaMongoDB(cadenaConexion, nombreBaseDatos)
	if err != nil {
		panic("Error al obtener la instancia de MongoDB:" + err.Error())
	}
	defer mongoDB.CerrarConexion()

	// Register MongoDB instance
	Container.Provide(func() *db.MongoDB {
		return mongoDB
	})

	// Register services
	Container.Provide(service.NewUserService)
	Container.Provide(service.NewChatService)

	// Register repositories
	Container.Provide(repository.NewUserRepository)
	Container.Provide(repository.NewChatRepository)

	// Register application
	Container.Provide(app.NewApplication)

	// Register handlers
	Container.Provide(handler.NewUserHandler)
	// Container.Provide(handler.NewChatHandler)

	return Container
}
