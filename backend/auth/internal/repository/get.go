package repository

import (
	"auth/internal/models"
	"errors"
	"go.uber.org/zap"
	"libs"
)

func (r AuthRepository) GetUser(email, password string) (models.UserDTO, error) {
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
		return models.UserDTO{}, queryError(log, op, err)
	}

	if !libs.CheckPass(user.Password, password) {
		return models.UserDTO{}, queryError(log, op, errors.New("Невалидные данные"))
	}

	return models.UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
