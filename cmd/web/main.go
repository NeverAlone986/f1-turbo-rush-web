package main

import (
	"log"
	"net/http"

	"github.com/yourusername/f1-turbo-rush/internal/game"
	"github.com/yourusername/f1-turbo-rush/internal/web"
)

func main() {
	// Инициализация игры
	game.InitGame()

	// Настройка маршрутов
	router := web.SetupRoutes()

	// Запуск сервера
	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
