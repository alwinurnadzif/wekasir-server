package utils

type ResponseWrapper struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
