package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *CloaksRepository) Get(ctx context.Context, id int) (string, error) {
	const op = "repository.Get"
	log := r.log.With(zap.String("method", op))

	query := `select path from cloaks where user_id = $1`

	var path string
	if err := r.db.
		QueryRowContext(
			ctx,
			query,
			id,
		).Scan(&path); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", libs.QueryError(log, op, err)
	}

	return path, nil
}
