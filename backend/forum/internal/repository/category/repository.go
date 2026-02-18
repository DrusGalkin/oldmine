package category

import (
	"context"
	"database/sql"
	"forum/internal/domain/models"
	"go.uber.org/zap"
)

type CategoryRepository interface {
	Get(ctx context.Context, id int) (models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
	GetByTemplateId(ctx context.Context, templateId int) ([]models.Category, error)
	Create(ctx context.Context, category models.Category) error
	Update(ctx context.Context, category models.Category) error
	Delete(ctx context.Context, id int) error
}

type CRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func New(db *sql.DB, log *zap.Logger) CategoryRepository {
	return &CRepository{
		db:  db,
		log: log,
	}
}
