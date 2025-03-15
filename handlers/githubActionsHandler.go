package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"webhook-github/config"

	"github.com/gin-gonic/gin"
)

// Estructura para recibir eventos de GitHub Actions
type GitHubActionsEvent struct {
	Workflow string `json:"workflow"`
	Action   string `json:"action"`
	Conclusion string `json:"conclusion"` // success, failure, etc.
}

// Manejador del webhook de GitHub Actions
func GitHubActionsWebhookHandler(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	var event GitHubActionsEvent
	err := json.Unmarshal(body, &event)
	if err != nil {
		log.Println("Error procesando el webhook de Actions:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Determinar el mensaje según el estado del workflow
	var message string
	switch event.Conclusion {
	case "success":
		message = fmt.Sprintf(" El workflow **%s** ha pasado exitosamente.", event.Workflow)
	case "failure":
		message = fmt.Sprintf(" El workflow **%s** ha fallado.", event.Workflow)
	default:
		message = fmt.Sprintf("ℹ Estado desconocido del workflow **%s**: %s", event.Workflow, event.Conclusion)
	}

	// Enviar notificación a Discord
	sendDiscordNotification(message, config.DiscordWebhookURLTests)

	c.JSON(http.StatusOK, gin.H{"message": "Webhook de GitHub Actions recibido"})
}
