package handler

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// Get godoc
// @Summary      Получить ссылку на путь к плащу
// @Description  ID пользователя = путь к файлу
// @Tags 		 Cloaks
// @Accept       json
// @Produce      json
// @Security	 CookieAuth
// @Param        id path int true "ID пользователя"
// @Success      200  {object}  object
// @Success      404  {object}  dto.ErrorResponse
// @Success		 500  {object}  dto.ErrorResponse
// @Router       /api/cloaks/{id} [get]
func (h *CloaksHandler) Get(ctx fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	path, err := h.repo.Get(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err,
			})
		} else {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"path": path,
	})
}
