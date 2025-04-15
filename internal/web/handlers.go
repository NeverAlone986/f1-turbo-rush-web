package web

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/f1-turbo-rush/internal/game"
)

func HandleGameState(w http.ResponseWriter, r *http.Request) {
	gameState := game.GetState()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameState)
}

func HandleTrackChange(w http.ResponseWriter, r *http.Request) {
	trackNum := r.URL.Query().Get("track")
	num, err := strconv.Atoi(trackNum)
	if err != nil || num < 1 || num > 5 { // Предполагаем 5 трасс
		http.Error(w, "Invalid track number", http.StatusBadRequest)
		return
	}
	
	game.ChangeTrack(num)
	w.WriteHeader(http.StatusOK)
}
