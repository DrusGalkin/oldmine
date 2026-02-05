package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"os"
	"path/filepath"
)

// Update godoc
// @Summary      Обновление скина
// @Description  Загрузка PNG файла скина
// @Tags         Skins
// @Accept       multipart/form-data
// @Produce      json
// @Security     CookieAuth
// @Param        file formData file true "PNG файл скина"
// @Success      202  {object}  object
// @Success      400  {object}  dto.ErrorResponse
// @Success      401  {object}  dto.ErrorResponse
// @Success      500  {object}  dto.ErrorResponse
// @Router       /api/skins/ [put]
func (h *SkinHandler) Update(ctx fiber.Ctx) error {
	formFile, err := ctx.FormFile(FORM_FILE_NAME)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	if filepath.Ext(formFile.Filename) != ".png" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Разрешенный формат только png",
		})
	}

	id := ctx.Locals("id").(int)
	fileName := ctx.Locals("name").(string) + ".png"
	filePath := fmt.Sprintf(
		"%s/%s",
		UPLOAD_PATH,
		fileName,
	)

	if err := h.repo.Update(ctx.Context(), id, filePath); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	os.Remove(filePath)
	if err := ctx.SaveFile(formFile, filePath); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{})
}
