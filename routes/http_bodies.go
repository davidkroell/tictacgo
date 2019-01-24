package routes

// Response is a default JSON respsonse body
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// NewGameBody is a JSON request body for a new game
type NewGameBody struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

// JoinGameBody is a JSON request body for joining a game
type JoinGameBody struct {
	Player string `json:"player"`
}

// PlayGameBody is a JSON request body for play game handler
type PlayGameBody struct {
	Player string `json:"player"`
	Field  int    `json:"fieldId"`
}
