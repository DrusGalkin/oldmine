package repository

import (
	"auth/internal/models"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"libs"
	"time"
)

func (r AuthRepository) Create(user models.User) error {
	const op = "repository.Create"
	log := r.log.With(zap.String("method", op))
	valid := validator.New()

	ctx, cancel := r.getContext()
	defer cancel()

	query := `insert into users(name, email, password, created_at) values ($1, $2, $3, $4) returning id`

	if err := valid.Struct(user); err != nil {
		return libs.QueryError(log, op, errors.New("Невалидные аргументы"))
	}

	var id int64
	user.Password = libs.HashPass(user.Password)
	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		time.Now(),
	).Scan(&id)
	if err != nil {
		return libs.QueryError(log, op, err)
	}

	if id == 0 {
		return libs.QueryError(log, op, errors.New("Пользователь не создан"))
	}

	return nil
}
