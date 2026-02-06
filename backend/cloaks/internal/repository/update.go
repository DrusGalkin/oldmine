package repository

import (
	"context"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *CloaksRepository) Update(ctx context.Context, uid int, path string) error {
	const op = "repository.Update"
	log := r.log.With(zap.String("method", op))

	query := `update cloaks set path=$1 where user_id=$2`

	_, err := r.db.ExecContext(
		ctx,
		query,
		path,
		uid,
	)
	if err != nil {
		return libs.QueryError(log, op, err)
	}

	return nil
}
