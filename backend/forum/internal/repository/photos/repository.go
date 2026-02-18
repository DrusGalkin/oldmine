package photos

import (
	"context"
	"database/sql"
	"forum/internal/domain/models"
	"go.uber.org/zap"
)

type PhotoRepository interface {
	Get(ctx context.Context, id int) (models.Photo, error)
	GetAll(ctx context.Context) ([]models.Photo, error)
	Create(ctx context.Context, photo models.Photo) error
	Update(ctx context.Context, id int, photo models.Photo) error
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
