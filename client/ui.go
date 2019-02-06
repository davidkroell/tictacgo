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

	joinUndefined := true
	var joingame bool

	for joinUndefined {
		fmt.Printf("Join [j] or create [c] game? ")
		input.Scan()
		join := input.Text()

		if len(join) == 0 {
			fmt.Print("Error: No input")
			continue
		}

		if join[0] != 'j' && join[0] != 'c' {
			fmt.Printf("Error occured:\nJoin or create a game!")
			continue
		}

		joingame = join[0] == 'j'
		joinUndefined = false
	}

	ch = make(chan string)

	if joingame {
		fmt.Printf("Game to join: ")
		input.Scan()
		go c.JoinGame(input.Text())

	} else {
		fmt.Printf("Name for new game: ")
		input.Scan()
		go c.CreateGame(input.Text(), ch)

		a := <-ch
		fmt.Print(a)
	}

}

func (c *Client) JoinGame(game string) {
	fmt.Printf("Joining game %s", game)
}
