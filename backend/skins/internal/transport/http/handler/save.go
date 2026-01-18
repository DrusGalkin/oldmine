package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"os"
	"path/filepath"
)

func (h *SkinHandler) Save(ctx fiber.Ctx) error {
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

	if err := ctx.SaveFile(formFile, filePath); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Скин установлен",
	})
}
