package photos

import (
	"context"
	"database/sql"
	"forum/internal/domain/model"
	"go.uber.org/zap"
)

type PhotoRepository interface {
	Get(ctx context.Context, id int) (model.Photo, error)
	GetAll(ctx context.Context) ([]model.Photo, error)
	Create(ctx context.Context, photo model.Photo) error
	Update(ctx context.Context, id int, photo model.Photo) error
	Delete(ctx context.Context, id int) error
}

type PRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func New(db *sql.DB, log *zap.Logger) PhotoRepository {
	return &PRepository{
		db:  db,
		log: log,
	}
}
