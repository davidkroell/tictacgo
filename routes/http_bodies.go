package routes

type Response struct {
	Success bool `json:"success"`
}

type SuccessResponse struct {
	Response
	Message string `json:"message"`
}

type ErrorResponse struct {
	SuccessResponse
	Code int `json:"code"`
}

type NewGameBody struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type JoinGameBody struct {
	Player string `json:"player"`
}

type NewGameResponse struct {
	SuccessResponse
}

type PlayGameBody struct {
	Player string `json:"player"`
	Field  int    `json:"fieldId"`
}
