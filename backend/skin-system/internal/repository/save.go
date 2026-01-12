package repository

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"libs"
)

func (r *SkinRepository) Save(ctx context.Context, uid int, path string) error {
	const op = "repository.Save"
	log := r.log.With(zap.String("method", op))

	query := "insert into skins(user_id, path) values ($1, $2) returning id"

	var id int
	if err := r.db.
		QueryRowContext(
			ctx,
			query,
			uid,
			path,
		).
		Scan(&id); err != nil {
		return libs.QueryError(log, op, err)
	}

	if id == 0 {
		return libs.QueryError(log, op, errors.New("id равен 0"))
	}

	return nil
}
