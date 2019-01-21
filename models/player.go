package models

import "github.com/gobuffalo/uuid"

type Player struct {
	ID   uuid.UUID
	Name string
}

func NewPlayer(name string) Player {
	id, _ := uuid.NewV4()

	return Player{
		ID:   id,
		Name: name,
	}
}
