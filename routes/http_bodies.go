package routes

// Default JSON respsonse body
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Request Body for a new game
type NewGameBody struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

// Request Body for joining a game
type JoinGameBody struct {
	Player string `json:"player"`
}

// Request Body for play game handler
type PlayGameBody struct {
	Player string `json:"player"`
	Field  int    `json:"fieldId"`
}
