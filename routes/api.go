package routes

import (
	"encoding/json"
	"fmt"
	"github.com/davidkroell/tictacgo/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Games stores all current games
var Games = map[string]models.Game{}

// NewGameHandler creates new game and save it into Games map
func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqBody NewGameBody

	err := decoder.Decode(&reqBody)
	if err != nil {
		log.Fatal(err.Error())
	}

	owner := models.NewPlayer(reqBody.Owner)
	Games[reqBody.Name] = models.NewGame(&owner)

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
	gameId := query["gameId"]

	game := Games[gameId]

	game.Player = models.NewPlayer(reqBody.Player)

	Games[gameId] = game

	var resp = Response{
		Success: true,
		Message: fmt.Sprintf("Game %s joined", query["gameId"]),
	}

	json.NewEncoder(w).Encode(resp)
}

// StatusGameHandler returns the status of a game, or an error if no game with this ID available
func StatusGameHandler(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	gameId := query["gameId"]

	game, exists := Games[gameId]
	if !exists {
		resp := Response{
			Success: false,
			Message: fmt.Sprintf("Game %s does not exist", gameId),
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
	gameId := query["gameId"]
	game := Games[gameId]

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
	Games[gameId] = game

	// write to response stream
	json.NewEncoder(w).Encode(game)
}
