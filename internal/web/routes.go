package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// API endpoints
	r.HandleFunc("/api/state", HandleGameState).Methods("GET")
	r.HandleFunc("/api/track/{number}", HandleChangeTrack).Methods("POST")

	// Static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	return r
}
