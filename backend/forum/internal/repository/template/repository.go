package template

import (
	"context"
	"database/sql"
	"forum/internal/domain/model"
	"go.uber.org/zap"
)

type TemplateRepository interface {
	Get(ctx context.Context, id int) (model.Template, error)
	GetAll(ctx context.Context) ([]model.Template, error)
	Create(ctx context.Context, tmp model.Template) error
	Update(ctx context.Context, id int, tmp model.Template) error
	Delete(ctx context.Context, id int) error
}

type TRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func New(db *sql.DB, log *zap.Logger) TemplateRepository {
	return &TRepository{
		db:  db,
		log: log,
	}
}
