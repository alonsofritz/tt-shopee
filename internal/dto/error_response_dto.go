package dto

type ErrorResponseDto struct {
	StatusCode  int    `json:"statusCode"`
	Error       string `json:"error"`
	Description string `json:"description"`
}
