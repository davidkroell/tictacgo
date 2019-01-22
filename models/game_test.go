package models

import (
	"testing"
)

func TestPlayGame(t *testing.T) {

	t.Run("Player wins first row", func(t *testing.T) {
		p1 := NewPlayer("David")
		p2 := NewPlayer("Christian")

		game := NewGame(&p1)
		game.JoinGame(&p2)

		err := game.PlayTurn(&game.Owner, 0)
		if err != nil {
			t.Errorf(err.Error() + " no error should occur")
		}

		err = game.PlayTurn(&p2, 3)
		if err != nil {
			t.Errorf(err.Error() + " no error should occur")
		}

		err = game.PlayTurn(&p1, 1)
		if err != nil {
			t.Errorf(err.Error() + " no error should occur")
		}

		err = game.PlayTurn(&game.Player, 4)
		if err != nil {
			t.Errorf(err.Error() + " no error should occur")
		}

		err = game.PlayTurn(&game.Owner, 2)
		if err != nil {
			t.Errorf(err.Error() + " no error should occur")
		}

		winner := game.Winner
		if winner == nil {
			t.Errorf(err.Error() + " " + p1.Name + " should win")
		}
	})

	t.Run("Player wins first coloumn", func(t *testing.T) {
		want := "Game already finished"
		p1 := NewPlayer("David")
		p2 := NewPlayer("Christian")

		game := NewGame(&p1)
		game.JoinGame(&p2)

		game.PlayTurn(&p1, 0)

		game.PlayTurn(&p2, 1)

		game.PlayTurn(&p1, 3)

		game.PlayTurn(&p2, 2)

		err := game.PlayTurn(&p1, 6)

		winner := game.Winner
		if winner == nil {
			t.Errorf(err.Error() + " " + p1.Name + " should win")
		}

		err = game.PlayTurn(&p2, 8)
		if err.Error() != want {
			t.Errorf("Got %s, want %s", err.Error(), want)
		}
	})

	t.Run("Player wins diagonal", func(t *testing.T) {
		p1 := NewPlayer("David")
		p2 := NewPlayer("Christian")

		game := NewGame(&p1)
		game.JoinGame(&p2)

		game.PlayTurn(&p1, 0)

		game.PlayTurn(&p2, 1)

		game.PlayTurn(&p1, 4)

		game.PlayTurn(&p2, 2)

		err := game.PlayTurn(&p1, 8)

		winner := game.Winner
		if winner == nil {
			t.Errorf(err.Error() + " " + p1.Name + " should win")
		}
	})

	t.Run("Player wins diagonal other side", func(t *testing.T) {
		p1 := NewPlayer("David")
		p2 := NewPlayer("Christian")

		game := NewGame(&p1)
		game.JoinGame(&p2)

		game.PlayTurn(&p1, 6)

		game.PlayTurn(&p2, 1)

		game.PlayTurn(&p1, 4)

		game.PlayTurn(&p2, 8)

		err := game.PlayTurn(&p1, 2)

		winner := game.Winner
		if winner == nil {
			t.Errorf(err.Error()+"%s should win", p1.Name)
		}
	})
}

func TestGameErrors(t *testing.T) {
	p1 := NewPlayer("David")
	p2 := NewPlayer("Christian")

	game := NewGame(&p1)
	game.JoinGame(&p2)

	t.Run("Same Field twice", func(t *testing.T) {
		want := "Field already occupied"

		err := game.PlayTurn(&game.Owner, 0)
		if err != nil {
			t.Errorf(err.Error() + " no error should occur")
		}

		err = game.PlayTurn(&game.Player, 0)
		if err.Error() != want {
			t.Errorf("Got %s, want %s", err.Error(), want)
		}
	})

	t.Run("Same Player twice", func(t *testing.T) {
		want := "Player not at turn"

		err := game.PlayTurn(&game.Owner, 3)
		if err != nil {
			t.Errorf(err.Error() + " no error should occur")
		}

		err = game.PlayTurn(&game.Owner, 4)
		if err.Error() != want {
			t.Errorf("Got %s, want %s", err.Error(), want)
		}
	})
}
