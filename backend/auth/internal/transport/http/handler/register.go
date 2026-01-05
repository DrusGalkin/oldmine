package handler

import (
	"auth/internal/models"
	"github.com/gofiber/fiber/v3"
)

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Невалидные данные: " + err.Error(),
		})
	}

	if err := h.repo.Create(models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при создании пользователя: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
}
