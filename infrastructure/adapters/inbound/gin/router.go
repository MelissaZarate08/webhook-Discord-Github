package gin

import (
	"github.com/gin-gonic/gin"
	"webhook-github/application"
	"webhook-github/infrastructure/adapters/outbound"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	devNotifier := discord.NewDiscordNotifier(false)
	testNotifier := discord.NewDiscordNotifier(true)

	prService := application.NewNotificationService(devNotifier)
	actionsService := application.NewNotificationService(testNotifier)

	r.POST("/webhook", NewGitHubWebhookHandler(prService))
	r.POST("/webhook-actions", NewGitHubActionsWebhookHandler(actionsService))

	return r
}
