package discord

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"webhook-github/infrastructure/config"
)

// Notifier define el comportamiento de cualquier notificador a Discord.
type Notifier interface {
	Send(message string) error
}

// DiscordNotifier implementa la interfaz Notifier.
type DiscordNotifier struct {
	webhookURL string
}

// NewDiscordNotifier crea un nuevo DiscordNotifier.
// Si isTest es true, se utiliza la URL de tests; en caso contrario, la de development.
func NewDiscordNotifier(isTest bool) *DiscordNotifier {
	url := config.DiscordWebhookURLDevelopment
	if isTest {
		url = config.DiscordWebhookURLTests
	}
	return &DiscordNotifier{
		webhookURL: url,
	}
}

// Send envía un mensaje al webhook configurado.
func (d *DiscordNotifier) Send(message string) error {
	data := map[string]string{"content": message}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error al serializar el mensaje:", err)
		return err
	}

	resp, err := http.Post(d.webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error enviando notificación a Discord:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		log.Printf("Error en la respuesta de Discord: %d\n", resp.StatusCode)
	}
	return nil
}