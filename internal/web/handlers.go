package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NeverAlone986/f1-turbo-rush-web/internal/game"
)

func HandleGameState(w http.ResponseWriter, r *http.Request) {
	currentState := game.GetState()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentState)
}

func HandleChangeTrack(w http.ResponseWriter, r *http.Request) {
	trackNum, err := strconv.Atoi(r.URL.Query().Get("track"))
	if err != nil {
		http.Error(w, "Invalid track number", http.StatusBadRequest)
		return
	}

	game.ChangeTrack(trackNum)
	w.WriteHeader(http.StatusOK)
}
