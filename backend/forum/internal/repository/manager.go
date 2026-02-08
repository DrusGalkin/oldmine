package repository

import (
	"database/sql"
	"forum/internal/repository/category"
	"forum/internal/repository/photos"
	"forum/internal/repository/template"
	"go.uber.org/zap"
)

type Repository struct {
	template.TemplateRepository
	category.CategoryRepository
	photos.PhotoRepository
}

func New(db *sql.DB, log *zap.Logger) Repository {
	return Repository{
		TemplateRepository: template.New(db, log),
		CategoryRepository: category.New(db, log),
		PhotoRepository:    photos.New(db, log),
	}
}
