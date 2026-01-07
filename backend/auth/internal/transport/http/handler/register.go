package handler

import (
	"auth/internal/dto"
	"auth/internal/models"
	"github.com/gofiber/fiber/v3"
)

// Register godoc
// @Summary      Регистрация
// @Description  Введите данные для регистрации
// @Tags 		 Auth
// @Accept       json
// @Produce      json
// @Param credential body dto.RegisterRequest true "Register credentials"
// @Success      201  {object}  object
// @Success		 400  {object}  dto.ErrorResponse
// @Success		 500  {object}  dto.ErrorResponse
// @Router       /register [post]
func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req dto.RegisterRequest
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
