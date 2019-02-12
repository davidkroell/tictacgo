package client

import (
	"encoding/json"
	"fmt"
	"github.com/davidkroell/tictacgo/models"
	"github.com/davidkroell/tictacgo/routes"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	healthRoute     = "/games"
	newGameRoute    = "/games/new"
	statusGameRoute = "/games/%s"
	joinGameRoute   = "/games/%s/join"
	playGameRoute   = "/games/%s/play"
)

func (c *Client) APIAlive(ch chan<- string) {
	resp, err := http.Get(c.BaseURL + healthRoute)

	if err != nil {
		ch <- fmt.Sprintf("Error occured in healthcheck:\n%v", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- fmt.Sprintf("Server not responding with HTTP 200. Aborting...")
		return
	}

	ch <- "alive"
}

// NewGame invokes a new game request to the api
func (c *Client) CreateGame(name string, ch chan<- string) {
	body := routes.NewGameBody{
		Name:  name,
		Owner: c.Username,
	}

	jsonresponse := jsonAPICall(c.BaseURL+newGameRoute, body, ch)

	if jsonresponse.Success {
		c.Game = name
	}

	ch <- jsonresponse.Message
}

func (c *Client) JoinGame(name string, ch chan<- string) {
	body := routes.JoinGameBody{
		Player: c.Username,
	}

	jsonresponse := jsonAPICall(fmt.Sprintf(c.BaseURL+joinGameRoute, name), body, ch)

	ch <- jsonresponse.Message
}

func jsonAPICall(endpoint string, body interface{}, ch chan<- string) routes.Response {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		if err := json.NewEncoder(writer).Encode(&body); err != nil {
			ch <- fmt.Sprintf("An error occured:\n%v", err.Error())
		}
	}()

	resp, err := http.Post(endpoint, "application/json", reader)
	if err != nil {
		ch <- fmt.Sprintf("An error occured:\n%v", err.Error())
		return routes.Response{}
	}
	defer resp.Body.Close()

	var jsonresponse routes.Response
	if err := json.NewDecoder(resp.Body).Decode(&jsonresponse); err != nil {
		ch <- fmt.Sprintf("An error occured:\n%v", err.Error())
		return routes.Response{}
	}

	return jsonresponse
}

func (c *Client) PlayTurn(field int, ch chan<- string) {
	body := routes.PlayGameBody{
		Player: c.Username,
		Field:  field,
	}

	jsonresponse := jsonAPICall(fmt.Sprintf(c.BaseURL+playGameRoute, c.Game), body, ch)

	if jsonresponse.Success {
		ch <- fmt.Sprintf("Turn %d played\n", field)
		return
	}

	ch <- jsonresponse.Message
}

func (c *Client) StatusGame() (models.Game, error) {
	resp, err := http.Get(fmt.Sprintf(c.BaseURL+statusGameRoute, c.Game))

	if err != nil {
		return models.Game{}, err
	}

	defer resp.Body.Close()

	var jsonresponse models.Game
	if err := json.NewDecoder(resp.Body).Decode(&jsonresponse); err != nil {
		return models.Game{}, err
	}

	return jsonresponse, nil
}

func (c *Client) StatusUpdater(interval time.Duration, ch chan models.Game) {
	for {
		game, err := c.StatusGame()
		if err != nil {
			log.Fatal(err)
		}

		select {
		case ch <- game:
		default:

		}
		time.Sleep(interval)
	}
}
