package handler

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v3"
	"os"
	"strconv"
)

// Delete godoc
// @Summary      Удаление скина
// @Description  Удаление скина пользователя
// @Tags         Skins
// @Accept       json
// @Produce      json
// @Security     CookieAuth
// @Param        id path int true "ID пользователя"
// @Success      204  {object}  object
// @Success      400  {object}  dto.ErrorResponse
// @Success      403  {object}  dto.ErrorResponse
// @Success      404  {object}  dto.ErrorResponse
// @Success      500  {object}  dto.ErrorResponse
// @Router       /api/skins/{id} [delete]
func (h *SkinHandler) Delete(ctx fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	localID := ctx.Locals("id").(int)
	admin := ctx.Locals("admin").(bool)
	name := ctx.Locals("name").(string)

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

	if admin {
		if err := h.repo.Delete(ctx.Context(), id); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		os.Remove(path)
	} else {

		if localID != id {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Нет доступа.",
			})
		}

		if err := h.repo.Delete(ctx.Context(), id); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		if err := os.Remove(path); err != nil {
			return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
				"message": "Ошибка удаления скина, возможно его не существует",
			})
		}
	}
	if path == "" {
		os.Remove(UPLOAD_PATH + "/" + name + ".png")
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Текущий скин удален",
	})
}
