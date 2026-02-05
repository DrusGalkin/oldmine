package handler

import (
	"auth/internal/dto"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

// Login godoc
// @Summary      Вход в аккаунт
// @Description  Введите логин и пароль
// @Tags 		 Auth
// @Accept       json
// @Produce      json
// @Param credential body dto.LoginRequest true "Login credentials"
// @Success      200  {object}  dto.User
// @Success		 400  {object}  dto.ErrorResponse
// @Success		 500  {object}  dto.ErrorResponse
// @Router       /api/auth/login [post]
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

	admin, _ := h.repo.IsAdmin(c.Context(), int64(user.ID))
	ver, _ := h.repo.PaymentVerification(c.Context(), int64(user.ID))

	sess.Set("id", user.ID)
	sess.Set("name", user.Name)
	sess.Set("email", user.Email)
	sess.Set("created_at", user.CreatedAt)
	sess.Set("auth", true)
	sess.Set("pay", ver)
	sess.Set("admin", admin)

	return c.JSON(fiber.Map{
		"session_id": sess.ID(),
		"user": dto.TrueLogResponse{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Admin:   admin,
			Payment: ver,
		},
	})
}
