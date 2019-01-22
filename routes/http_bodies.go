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

type NewGameResponse struct {
	SuccessResponse
}
