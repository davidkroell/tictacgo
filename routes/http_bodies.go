package routes

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type NewGameBody struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type JoinGameBody struct {
	Player string `json:"player"`
}

type PlayGameBody struct {
	Player string `json:"player"`
	Field  int    `json:"fieldId"`
}
