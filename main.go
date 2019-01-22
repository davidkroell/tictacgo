package main

import (
	. "davidkroell/basichttp/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	gamesRouter := r.PathPrefix("/games").Subrouter()

	gamesRouter.HandleFunc("/new", NewGameHandler).Methods("POST")
	gamesRouter.HandleFunc("/{gameId}", StatusGameHandler).Methods("GET")
	gamesRouter.HandleFunc("/{gameId}/join", JoinGameHandler).Methods("POST")
	gamesRouter.HandleFunc("/{gameId}/play", PlayGameHandler).Methods("POST")

	gamesRouter.Use(RequestLogger)
	gamesRouter.Use(HeaderBinding)

	log.Fatal(http.ListenAndServe(":9999", r))
}
