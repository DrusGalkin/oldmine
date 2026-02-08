package category

import (
	"context"
	"database/sql"
	"forum/internal/domain/model"
	"go.uber.org/zap"
)

type CategoryRepository interface {
	Get(ctx context.Context, id int) (model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)
	GetByTemplateId(ctx context.Context, templateId int) ([]model.Category, error)
	Create(ctx context.Context, category model.Category) error
	Update(ctx context.Context, category model.Category) error
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
