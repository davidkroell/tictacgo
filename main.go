package main

import (
	"github.com/davidkroell/tictacgo/client"
	"github.com/davidkroell/tictacgo/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		ClientMode()
	} else {
		ServerMode()
	}
}

// ServerMode starts a tictacgo server instance
func ServerMode() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file present. Environment not loaded from file")
	}

	// init router
	r := mux.NewRouter()
	gamesRouter := r.PathPrefix("/games").Subrouter()

	// bind routes
	gamesRouter.HandleFunc("", routes.HealthHandler).Methods("GET")
	gamesRouter.HandleFunc("/new", routes.NewGameHandler).Methods("POST")
	gamesRouter.HandleFunc("/{gameID}", routes.StatusGameHandler).Methods("GET")
	gamesRouter.HandleFunc("/{gameID}/join", routes.JoinGameHandler).Methods("POST")
	gamesRouter.HandleFunc("/{gameID}/play", routes.PlayGameHandler).Methods("POST")

	// add middleware
	gamesRouter.Use(routes.RequestLogger)
	gamesRouter.Use(routes.HeaderBinding)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port environment variable not set")
	}

	// start server
	log.Printf("Starting server on %s:%s", os.Getenv("LISTEN_ADDR"), port)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDR")+":"+port, r))
}

// ClientMode launches the client application
func ClientMode() {
	url := os.Args[1]

	// add https if no protocol is specified
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	// cut trailing comma
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	c := client.NewClient(os.Args[1])

	c.StartInteractive()
}
