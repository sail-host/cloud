package dto

type BaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type BaseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
