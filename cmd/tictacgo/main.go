package main

import (
	"github.com/davidkroell/tictacgo"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		tictacgo.ClientMode()
	} else {
		tictacgo.ServerMode()
	}
}
