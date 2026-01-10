package repository

import (
	"context"
	"go.uber.org/zap"
)

func (r AuthRepository) PaymentVerification(ctx context.Context, reqID int64) (bool, error) {
	const op = "repository.PaymentVerification"
	log := r.log.With(zap.String("method", op))

	query := `select user_id from paid_for where user_id = $1`

	var userID int64
	if err := r.db.QueryRowContext(
		ctx,
		query,
		reqID,
	).Scan(&userID); err != nil {
		return false, queryError(log, op, err)
	}

	return userID == reqID, nil
}
