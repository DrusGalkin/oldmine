package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r AuthRepository) IsAdmin(ctx context.Context, reqID int64) (bool, error) {
	const op = "repository.IsAdmin"
	log := r.log.With(zap.String("method", op))

	query := `select user_id from admins where user_id = $1`

	var userID int64
	if err := r.db.QueryRowContext(
		ctx,
		query,
		reqID,
	).Scan(&userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, libs.QueryError(log, op, err)
	}

	return userID == reqID, nil
}
