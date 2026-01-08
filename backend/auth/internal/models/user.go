package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=16"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8,max=32"`
	CreatedAt time.Time `json:"created_at"`
}
