package repository

import (
	"auth/internal/dto"
	"auth/internal/models"
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type Repository interface {
	Create(user models.User) error
	GetUser(email, password string) (dto.User, error)
}

type AuthRepository struct {
	db  *sql.DB
	log *zap.Logger
	ttl time.Duration
}

func New(db *sql.DB, log *zap.Logger, ttl time.Duration) Repository {
	if ttl == 0 {
		ttl = 10 * time.Second
	}

	return AuthRepository{
		db:  db,
		log: log,
		ttl: ttl,
	}
}

func (r AuthRepository) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), r.ttl)
}

func queryError(log *zap.Logger, op string, err error) error {
	msg := fmt.Sprintf("%s: %s", op, err.Error())
	log.Error(msg)
	return err
}
