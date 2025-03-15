package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Enviar mensaje a Discord
func sendDiscordNotification(message, webhookURL string) {
	data := map[string]string{"content": message}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error enviando notificación a Discord:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Printf("Error en la respuesta de Discord: %d\n", resp.StatusCode)
	}
}
