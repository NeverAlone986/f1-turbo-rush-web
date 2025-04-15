package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	
	// API routes
	r.HandleFunc("/api/game-state", HandleGameState).Methods("GET")
	r.HandleFunc("/api/change-track/{track}", HandleTrackChange).Methods("POST")
	
	// Static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	
	return r
}
