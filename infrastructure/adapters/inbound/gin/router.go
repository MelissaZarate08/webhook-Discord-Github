package gin

import (
	"github.com/gin-gonic/gin"
	"webhook-github/application"
	"webhook-github/infrastructure/adapters/outbound"
)

// NewRouter configura y retorna el router con las rutas necesarias.
func NewRouter() *gin.Engine {
	r := gin.Default()

	// Se crean instancias del notificador según el contexto.
	devNotifier := discord.NewDiscordNotifier(false)
	testNotifier := discord.NewDiscordNotifier(true)

	// Se crean los servicios inyectando el notificador adecuado.
	prService := application.NewNotificationService(devNotifier)
	actionsService := application.NewNotificationService(testNotifier)

	// Rutas para los distintos webhooks.
	r.POST("/webhook", NewGitHubWebhookHandler(prService))
	r.POST("/webhook-actions", NewGitHubActionsWebhookHandler(actionsService))

	return r
}
