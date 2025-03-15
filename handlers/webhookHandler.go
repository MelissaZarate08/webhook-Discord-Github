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

// Estructura para recibir eventos de GitHub
type GitHubEvent struct {
	Action      string `json:"action"`
	PullRequest struct {
		Title string `json:"title"`
		Number int `json:"number"`
		User struct {
			Login string `json:"login"`
		} `json:"user"`
	} `json:"pull_request"`
}

// Manejador del webhook de GitHub
func GitHubWebhookHandler(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	var event GitHubEvent
	err := json.Unmarshal(body, &event)
	if err != nil {
		log.Println("Error procesando el webhook:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Construcción del mensaje
	var message string
	switch event.Action {
	case "opened":
		message = fmt.Sprintf("Nuevo PR: **%s** (#%d) abierto por @%s", event.PullRequest.Title, event.PullRequest.Number, event.PullRequest.User.Login)
	case "reopened":
		message = fmt.Sprintf("PR reabierto: **%s** (#%d)", event.PullRequest.Title, event.PullRequest.Number)
	case "ready_for_review":
		message = fmt.Sprintf("PR listo para revisión: **%s** (#%d)", event.PullRequest.Title, event.PullRequest.Number)
	case "closed":
		message = fmt.Sprintf("PR cerrado o fusionado: **%s** (#%d)", event.PullRequest.Title, event.PullRequest.Number)
	}

	// Enviar notificación a Discord si hay un mensaje
	if message != "" {
		sendDiscordNotification(message, config.DiscordWebhookURLDevelopment)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook recibido"})
}
