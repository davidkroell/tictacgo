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

	// init router
	r := mux.NewRouter()
	gamesRouter := r.PathPrefix("/games").Subrouter()

	// bind routes
	gamesRouter.HandleFunc("/new", NewGameHandler).Methods("POST")
	gamesRouter.HandleFunc("/{gameId}", StatusGameHandler).Methods("GET")
	gamesRouter.HandleFunc("/{gameId}/join", JoinGameHandler).Methods("POST")
	gamesRouter.HandleFunc("/{gameId}/play", PlayGameHandler).Methods("POST")

	// add middleware
	gamesRouter.Use(RequestLogger)
	gamesRouter.Use(HeaderBinding)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port environment variable not set")
	}

	// start server
	log.Printf("Starting server on %s:%s", os.Getenv("LISTEN_ADDR"), port)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDR")+":"+port, r))
}
