package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/oliamb/cutter"
	"image/png"
	"os"
	"path/filepath"
)

// Save godoc
// @Summary      Загрузка скина
// @Description  Загрузка PNG файла скина
// @Tags         Skins
// @Accept       multipart/form-data
// @Produce      json
// @Security     CookieAuth
// @Param        file formData file true "PNG файл скина"
// @Success      201  {object}  object
// @Success      400  {object}  dto.ErrorResponse
// @Success      401  {object}  dto.ErrorResponse
// @Success      500  {object}  dto.ErrorResponse
// @Router       /api/skins/ [post]
func (h *SkinHandler) Save(ctx fiber.Ctx) error {
	formFile, err := ctx.FormFile(FORM_FILE_NAME)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if filepath.Ext(formFile.Filename) != ".png" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Разрешенный формат только png",
		})
	}

	tmpFile, err := os.CreateTemp("", "upload_skin_*.png")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось создать временный файл",
		})
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	if err := ctx.SaveFile(formFile, tmpFile.Name()); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось сохранить загруженный файл",
		})
	}

	file, err := os.Open(tmpFile.Name())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось открыть файл изображения",
		})
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Некорректный PNG файл",
		})
	}

	bounds := img.Bounds()
	if bounds.Dx() > 64 || bounds.Dy() > 64 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Изображение должно быть размером 64x64 или 64х32 пикселей",
		})
	}

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  64,
		Height: 32,
		Mode:   cutter.TopLeft,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось обрезать изображение",
		})
	}

	id := ctx.Locals("id").(int)
	fileName := ctx.Locals("name").(string) + ".png"
	filePath := fmt.Sprintf("%s/%s", UPLOAD_PATH, fileName)

	// Возможно, кто-то захочет сохранить скин, когда он уже существует, значит нужно понять, хранится ли в бд
	// тот же путь на файл, и да, возможно метод update не нужен, но для семантики пусть будет.

	var flag bool
	path, _ := h.repo.Get(ctx.Context(), id)
	if len(path) != 0 {
		if path == filePath {
			os.Remove(filePath)
			flag = true // если данные пользователя уже есть в бд, то устанавливаем флаг на true
		}
	}

	// данных нет в бд, добавляем
	if !flag {
		if err := h.repo.Save(ctx.Context(), id, filePath); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось создать файл для сохранения",
		})
	}

	defer outputFile.Close()

	if err := png.Encode(outputFile, croppedImg); err != nil {
		os.Remove(filePath)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось сохранить обрезанное изображение",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Скин установлен",
	})
}
