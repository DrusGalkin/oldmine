package repository

import (
	"context"
	"go.uber.org/zap"
	"libs"
)

func (r *SkinRepository) Delete(ctx context.Context, id int) error {
	const op = "repository.Delete"
	log := r.log.With(zap.String("method", op))

	query := `delete from skins where id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return libs.QueryError(log, op, err)
	}

	return nil
}
