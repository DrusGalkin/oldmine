package dto

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Payment   bool      `json:"payment"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"created_at"`
}
