package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *CloaksRepository) Delete(ctx context.Context, id int) error {
	const op = "repository.Delete"
	log := r.log.With(zap.String("method", op))

	query := `delete from cloaks where user_id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return libs.QueryError(log, op, err)
	}

	return nil
}
