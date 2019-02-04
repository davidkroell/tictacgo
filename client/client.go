package client

import (
	"bufio"
	"fmt"
	"os"
)

type Client struct {
	BaseURL  string
	Username string
}

func NewClient(baseURL string) Client {
	return Client{
		BaseURL: baseURL,
	}
}

func (c *Client) StartInteractive() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Username: ")
	input.Scan()
	c.Username = input.Text()

	joinUndefined := true
	var joingame bool

	for joinUndefined {
		fmt.Printf("Join [j] or create [c] game? ")
		input.Scan()
		join := input.Text()

		if join[0] != 'j' && join[0] != 'c' {
			fmt.Printf("[  ERR  ] Join or create a game!")
			continue
		}

		joingame = join[0] == 'j'
		joinUndefined = false
	}

	if joingame {
		fmt.Printf("Name for new game: ")
		input.Scan()
		c.JoinGame(input.Text())
	}

	fmt.Printf("Game to join: ")
	input.Scan()
	c.CreateGame(input.Text())
}

func (c *Client) JoinGame(game string) {
	fmt.Printf("Joining game %s", game)
}

func (c *Client) CreateGame(game string) {
	fmt.Printf("Creating game %s", game)
}
