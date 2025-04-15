package main

import (
	"log"
	"net/http"

	"github.com/NeverAlone986/f1-turbo-rush-web/internal/game"
	"github.com/NeverAlone986/f1-turbo-rush-web/internal/web"
)

func main() {
	game.InitGame()

	// Инициализация HTTP сервера
	router := web.SetupRoutes()

	// Запуск сервера
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
