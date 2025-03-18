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

// GitHubActionsEvent representa el payload del webhook de GitHub Actions.
type GitHubActionsEvent struct {
	Action      string `json:"action"`
	WorkflowRun struct {
		Name       string `json:"name"`
		Conclusion string `json:"conclusion"`
	} `json:"workflow_run"`
}

// NewGitHubActionsWebhookHandler crea un handler para los eventos de GitHub Actions.
func NewGitHubActionsWebhookHandler(svc *application.NotificationService) gin.HandlerFunc {
    return func(c *gin.Context) {
        body, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            log.Println("Error al leer el cuerpo:", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo leer el cuerpo"})
            return
        }

        // Imprimir el cuerpo completo del payload para depuración
        log.Println("Payload recibido:", string(body))

        var event GitHubActionsEvent
        if err := json.Unmarshal(body, &event); err != nil {
            log.Println("Error procesando el webhook de Actions:", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
            return
        }

        log.Println("Evento procesado:", event)  // Verifica que los valores son correctos

        actionsEvent := domain.ActionsEvent{
            Workflow:   event.WorkflowRun.Name,       // Extraemos el nombre del workflow
            Action:     event.Action,                  // Este valor ya lo tienes
            Conclusion: event.WorkflowRun.Conclusion,  // Extraemos la conclusión
        }

        if err := svc.NotifyActionsEvent(actionsEvent); err != nil {
            log.Println("Error notificando a Discord:", err)
        }

        c.JSON(http.StatusOK, gin.H{"message": "Webhook de Actions recibido"})
    }
}