package repository

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
)

type Repository interface {
	Save(ctx context.Context, uid int, path string) error
	Get(ctx context.Context, id int) (string, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, uid int, path string) error
}

type SkinRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func New(db *sql.DB, log *zap.Logger) Repository {
	return &SkinRepository{
		db:  db,
		log: log,
	}
}
