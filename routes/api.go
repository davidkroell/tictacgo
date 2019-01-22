package routes

import (
	"davidkroell/basichttp/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Games = map[string]models.Game{}

// Create new game and save into Games map
func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqBody NewGameBody

	err := decoder.Decode(&reqBody)
	if err != nil {
		log.Fatal(err.Error())
	}

	owner := models.NewPlayer(reqBody.Owner)
	Games[reqBody.Name] = models.NewGame(&owner)

	var resp = NewGameResponse{
		SuccessResponse: SuccessResponse{
			Response: Response{
				Success: true,
			},
			Message: fmt.Sprintf("Game %s created", reqBody.Name),
		},
	}

	json.NewEncoder(w).Encode(resp)
}

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

	var resp = NewGameResponse{
		SuccessResponse: SuccessResponse{
			Response: Response{
				Success: true,
			},
			Message: fmt.Sprintf("Game %s joined", query["gameId"]),
		},
	}

	json.NewEncoder(w).Encode(resp)
}

func StatusGameHandler(w http.ResponseWriter, r *http.Request) {

}

func PlayGameHandler(w http.ResponseWriter, r *http.Request) {

}
