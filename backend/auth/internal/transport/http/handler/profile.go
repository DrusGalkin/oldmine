package handler

import (
	"auth/internal/dto"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"time"
)

// Profile godoc
// @Summary      Профиль текущей сессии
// @Description  Для того чтобы получить данные пользователя, требуется авторизация
// @Tags 		 Auth
// @Accept       json
// @Produce      json
// @Security	 CookieAuth
// @Success      200  {object}	dto.User
// @Success		 401  {object}  object
// @Router       /api/auth/profile [get]
func (h *AuthHandler) Profile(c fiber.Ctx) error {
	sess := session.FromContext(c)
	return c.JSON(dto.User{
		ID:        sess.Get("id").(int),
		Name:      sess.Get("name").(string),
		Email:     sess.Get("email").(string),
		Payment:   sess.Get("pay").(bool),
		Admin:     sess.Get("admin").(bool),
		CreatedAt: sess.Get("created_at").(time.Time),
	})
}
