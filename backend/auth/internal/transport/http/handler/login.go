package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Невалидные данные: " + err.Error(),
		})
	}

	user, err := h.repo.GetUser(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	sess := session.FromContext(c)

	if err := sess.Regenerate(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	sess.Set("auth", true)
	sess.Set("id", user.ID)
	sess.Set("name", user.Name)
	sess.Set("email", user.Email)

	return c.JSON(fiber.Map{
		"session_id": sess.ID(),
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
