package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TrueLogResponse struct {
	ID      int
	Name    string
	Email   string
	Admin   bool
	Payment bool
}

type LoginResponse struct {
	SessID string `json:"sessID"`
	User   TrueLogResponse
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
