// container.go
package di

import (
	"github.com/luisk6510/ChatEnTiempoReal/internal/app"
	"github.com/luisk6510/ChatEnTiempoReal/internal/repository"
	"github.com/luisk6510/ChatEnTiempoReal/internal/service"
	"github.com/luisk6510/ChatEnTiempoReal/web/handler"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Register services
	container.Provide(service.NewUserService)
	container.Provide(service.NewChatService)

	// Register repositories
	container.Provide(repository.NewUserRepository)
	container.Provide(repository.NewChatRepository)

	// Register application
	container.Provide(app.NewApplication)

	// Register handlers
	container.Provide(handler.NewUserHandler)
	// container.Provide(handler.NewChatHandler)

	return container
}
