package main

import (
	"log"
	"webhook-github/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ruta para recibir los webhooks de GitHub PRs
	r.POST("/webhook", handlers.GitHubWebhookHandler)

	// Ruta para recibir los webhooks de GitHub Actions
	r.POST("/webhook-actions", handlers.GitHubActionsWebhookHandler)

	// Iniciar el servidor en el puerto 8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
