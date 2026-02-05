package repository

import (
	"auth/internal/dto"
	"auth/internal/models"
	"errors"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r AuthRepository) GetUser(email, password string) (dto.User, error) {
	const op = "repository.GetUser"
	log := r.log.With(zap.String("method", op))

	ctx, cancel := r.getContext()
	defer cancel()

	query := `select id, name, email, password, created_at from users where email = $1`

	user := models.User{}
	if err := r.db.QueryRowContext(
		ctx,
		query,
		email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		return dto.User{}, libs.QueryError(log, op, err)
	}

	if !libs.CheckPass(user.Password, password) {
		return dto.User{}, libs.QueryError(log, op, errors.New("Невалидные данные"))
	}

	return dto.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
