package main

import (
	"log"

	"webhook-github/infrastructure/adapters/inbound/gin"
)

func main() {
	router := gin.NewRouter()

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
