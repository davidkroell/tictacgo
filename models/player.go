package models

type Player struct {
	Name string `json:"name"`
}

func NewPlayer(name string) Player {

	return Player{
		Name: name,
	}
}
