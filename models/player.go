package models

//Player hold data about a player
type Player struct {
	Name string `json:"name"`
}

//NewPlayer returns a new instance of a Player
func NewPlayer(name string) Player {

	return Player{
		Name: name,
	}
}
