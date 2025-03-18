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


type GitHubActionsEvent struct {
    Workflow json.RawMessage `json:"workflow"`
    WorkflowRun struct {
        Name        string `json:"name"`
        Conclusion  string `json:"conclusion"`
    } `json:"workflow_run"`
    Action string `json:"action"`
}



func NewGitHubActionsWebhookHandler(svc *application.NotificationService) gin.HandlerFunc {
    return func(c *gin.Context) {
        body, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            log.Println("Error al leer el cuerpo:", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo leer el cuerpo"})
            return
        }


        var event GitHubActionsEvent
        if err := json.Unmarshal(body, &event); err != nil {
            log.Println("Error procesando el webhook de Actions:", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
            return
        }

		
        if event.WorkflowRun.Name == "" || event.WorkflowRun.Conclusion == "" {
            log.Println("Evento ignorado por falta de datos en workflow_run")
            c.JSON(http.StatusOK, gin.H{"message": "Evento ignorado"})
            return
        }


        actionsEvent := domain.ActionsEvent{
            Workflow:   event.WorkflowRun.Name,
            Action:     event.Action,
            Conclusion: event.WorkflowRun.Conclusion,
        }

        if err := svc.NotifyActionsEvent(actionsEvent); err != nil {
            log.Println("Error notificando a Discord:", err)
        }

        c.JSON(http.StatusOK, gin.H{"message": "Webhook de Actions recibido"})
    }
}
