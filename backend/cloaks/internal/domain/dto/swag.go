package dto

type ErrorResponse struct {
	Error string `json:"error"`
}

type SkinRequest struct {
	ID string `json:"id"`
}
