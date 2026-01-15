package repository

import (
	"context"
	"go.uber.org/zap"
	"libs"
)

func (r *SkinRepository) Get(ctx context.Context, id int) (string, error) {
	const op = "repository.Get"
	log := r.log.With(zap.String("method", op))

	query := `select path from skins where user_id = $1`

	var path string
	if err := r.db.
		QueryRowContext(
			ctx,
			query,
			id,
		).Scan(&path); err != nil {
		return "", libs.QueryError(log, op, err)
	}

	return path, nil
}
