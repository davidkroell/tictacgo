package client

import (
	"encoding/json"
	"fmt"
	"github.com/davidkroell/tictacgo/routes"
	"io"
	"net/http"
)

const (
	prefix = "/games"

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

	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		if err := json.NewEncoder(writer).Encode(&body); err != nil {
			ch <- fmt.Sprintf("An error occured:\n%v", err.Error())
		}
	}()

	resp, err := http.Post(c.BaseURL+newGameRoute, "application/json", reader)
	if err != nil {
		ch <- fmt.Sprintf("An error occured:\n%v", err.Error())
		return
	}
	defer resp.Body.Close()

	var jsonresponse routes.Response
	if err := json.NewDecoder(resp.Body).Decode(&jsonresponse); err != nil {
		ch <- fmt.Sprintf("An error occured:\n%v", err.Error())
		return
	}

	if jsonresponse.Success {
		ch <- fmt.Sprintf("Game %s created", name)
		return
	}

	ch <- fmt.Sprint("Error occured")
}
