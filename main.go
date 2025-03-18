package main

import (
	"log"

	"webhook-github/infrastructure/adapters/inbound/gin"
)

func main() {
	// Crea el router configurado con los handlers
	router := gin.NewRouter()

	// Iniciar el servidor en el puerto 8080
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}