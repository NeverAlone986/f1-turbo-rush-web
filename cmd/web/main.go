package main

import (
	"log"
	"net/http"

	"github.com/NeverAlone986/f1-turbo-rush-web/internal/game"
	"github.com/NeverAlone986/f1-turbo-rush-web/internal/web"
)

func main() {
	game.InitGame() // Теперь функция доступна
	router := web.SetupRoutes()
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
