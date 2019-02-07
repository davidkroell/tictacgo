package client

import (
	"bufio"
	"fmt"
	"github.com/davidkroell/tictacgo/models"
	"os"
	"strconv"
	"time"
)

// Client is a
type Client struct {
	BaseURL  string
	Username string
	Game     string
}

// NewClient returns a new client struct
func NewClient(baseURL string) Client {
	return Client{
		BaseURL: baseURL,
	}
}

// StartInteractive does a basic setup and starts the client interactively
func (c *Client) StartInteractive() {
	ch := make(chan string)

	go c.APIAlive(ch)
	if alive := <-ch; alive != "alive" {
		fmt.Print(alive)
		return
	}

	input := bufio.NewScanner(os.Stdin)
	defer os.Stdin.Close()

	fmt.Printf("Username: ")
	input.Scan()
	c.Username = input.Text()

	joingame := isJoinGame(input)

	ch = make(chan string)

	if joingame {
		fmt.Printf("Game to join: ")
		input.Scan()
		go c.JoinGame(input.Text(), ch)

		fmt.Print(<-ch)
	} else {
		fmt.Printf("Name for new game: ")
		input.Scan()
		go c.CreateGame(input.Text(), ch)

		fmt.Print(<-ch)
	}
	c.GameLoop()
}

func isJoinGame(input *bufio.Scanner) bool {
	joinUndefined := true

	for joinUndefined {
		fmt.Printf("Join [j] or create [c] game? ")
		input.Scan()
		join := input.Text()

		if len(join) == 0 {
			fmt.Println("Error: No input")
			continue
		}

		if join[0] != 'j' && join[0] != 'c' {
			fmt.Printf("Error occured: Join or create a game!\n")
			continue
		}

		return join[0] == 'j'
		joinUndefined = false
	}

	return false
}

// GameLoop starts the main game loop until the game is finished
func (c *Client) GameLoop() {
	// TODO complete

	uich := make(chan models.Game)
	ms := 1000
	interval := time.Duration(ms) * time.Millisecond
	go c.StatusUpdater(interval, uich)

	ch := make(chan string)

	input := bufio.NewScanner(os.Stdin)
	for {
		gamestatus := <-uich

		if gamestatus.Player == (models.Player{}) {
			fmt.Println("Waiting for other player")
			continue
		}

		err := c.RenderPlayField(gamestatus, os.Stdout)

		if err != nil {
			fmt.Print(err)
		}

		if gamestatus.NextTurn.Name == c.Username {
			// current players turn
			fmt.Print("Field number: ")

			input.Scan()
			num, err := strconv.Atoi(input.Text())
			if err != nil {
				fmt.Println("Error occured. Type a number")
				continue
			}
			go c.PlayTurn(num, ch)
			fmt.Print(<-ch) //TODO implement
		}
	}
}
