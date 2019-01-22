package main

import (
	. "github.com/davidkroell/tictacgo/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file present. Environment not loaded from file")
	}

	r := mux.NewRouter()
	gamesRouter := r.PathPrefix("/games").Subrouter()

	gamesRouter.HandleFunc("/new", NewGameHandler).Methods("POST")
	gamesRouter.HandleFunc("/{gameId}", StatusGameHandler).Methods("GET")
	gamesRouter.HandleFunc("/{gameId}/join", JoinGameHandler).Methods("POST")
	gamesRouter.HandleFunc("/{gameId}/play", PlayGameHandler).Methods("POST")

	gamesRouter.Use(RequestLogger)
	gamesRouter.Use(HeaderBinding)

	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDR")+":"+os.Getenv("PORT"), r))
}
