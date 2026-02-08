package repository

import (
	"forum/internal/repository/category"
	"forum/internal/repository/photo"
	"forum/internal/repository/tempalte"
)

type Repository struct {
	tempalte.TemplateRepository
	category.CategoryRepository
	photo.PhotoRepository
}
