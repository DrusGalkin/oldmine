package handler

import (
	"auth/internal/repository"
	"github.com/gofiber/fiber/v3"
)

type Handler interface {
	Login(c fiber.Ctx) error
	Profile(c fiber.Ctx) error
	Register(c fiber.Ctx) error
	Logout(c fiber.Ctx) error
}

type AuthHandler struct {
	repo repository.Repository
}

func New(repo repository.Repository) Handler {
	return &AuthHandler{repo: repo}
}
