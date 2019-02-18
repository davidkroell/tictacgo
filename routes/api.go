package routes

import (
	"encoding/json"
	"fmt"
	"github.com/davidkroell/tictacgo/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

// Games stores all current games
var Games = struct {
	Collection map[string]models.Game
	sync.Mutex
}{
	Collection: map[string]models.Game{},
}

// HealthHandler returns HTTP 200 if service is available
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"healthy"}`))
}

// NewGameHandler creates new game and save it into Games map
func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqBody NewGameBody

	if err := decoder.Decode(&reqBody); err != nil {
		log.Fatal(err.Error())
	}

	Games.Lock()
	defer Games.Unlock()

	if _, exists := Games.Collection[reqBody.Name]; exists {
		var resp = Response{
			Success: false,
			Message: fmt.Sprintf("Game named %s already exists", reqBody.Name),
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	owner := models.NewPlayer(reqBody.Owner)
	Games.Collection[reqBody.Name] = models.NewGame(&owner)

	var resp = Response{
		Success: true,
		Message: fmt.Sprintf("Game %s created", reqBody.Name),
	}

	json.NewEncoder(w).Encode(resp)
}

// JoinGameHandler handles join of a player
func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqBody JoinGameBody

	err := decoder.Decode(&reqBody)
	if err != nil {
		log.Fatal(err.Error())
	}
	query := mux.Vars(r)
	gameID := query["gameID"]

	Games.Lock()
	game := Games.Collection[gameID]

	game.Player = models.NewPlayer(reqBody.Player)

	Games.Collection[gameID] = game
	Games.Unlock()

	var resp = Response{
		Success: true,
		Message: fmt.Sprintf("Game %s joined", query["gameID"]),
	}

	json.NewEncoder(w).Encode(resp)
}

// StatusGameHandler returns the status of a game, or an error if no game with this ID available
func StatusGameHandler(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	gameID := query["gameID"]

	game, exists := Games.Collection[gameID]
	if !exists {
		resp := Response{
			Success: false,
			Message: fmt.Sprintf("Game %s does not exist", gameID),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(game)
}

// PlayGameHandler handles playing a game with the given Player and field ID
func PlayGameHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve game
	query := mux.Vars(r)
	gameID := query["gameID"]

	Games.Lock()
	defer Games.Unlock()
	game := Games.Collection[gameID]

	// retrieve json data
	decoder := json.NewDecoder(r.Body)
	var reqBody PlayGameBody
	err := decoder.Decode(&reqBody)
	if err != nil {
		log.Fatal(err.Error())
	}

	// get player pointer
	var player models.Player
	if game.Player.Name == reqBody.Player {
		player = game.Player
	} else if game.Owner.Name == reqBody.Player {
		player = game.Owner
	}

	err = game.PlayTurn(&player, reqBody.Field)
	if err != nil {
		resp := Response{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	// save game back
	Games.Collection[gameID] = game

	// write to response stream
	json.NewEncoder(w).Encode(game)
}
