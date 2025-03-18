package gin

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"webhook-github/application"
    "webhook-github/domain"
)

// GitHubEvent representa el payload del webhook de un Pull Request.
type GitHubEvent struct {
	Action      string `json:"action"`
	PullRequest struct {
		Title  string `json:"title"`
		Number int    `json:"number"`
		User   struct {
			Login string `json:"login"`
		} `json:"user"`
	} `json:"pull_request"`
}

// NewGitHubWebhookHandler crea un handler para los eventos de PR.
func NewGitHubWebhookHandler(svc *application.NotificationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Error al leer el cuerpo:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo leer el cuerpo"})
			return
		}

		var event GitHubEvent
		if err := json.Unmarshal(body, &event); err != nil {
			log.Println("Error procesando el webhook:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
			return
		}

		prEvent := domain.PullRequestEvent{
			Title:  event.PullRequest.Title,
			Number: event.PullRequest.Number,
			User:   event.PullRequest.User.Login,
			Action: event.Action,
		}

		if err := svc.NotifyPullRequestEvent(prEvent); err != nil {
			log.Println("Error notificando a Discord:", err)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Webhook recibido"})
	}
}
