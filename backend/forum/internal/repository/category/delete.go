package category

import (
	"context"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *CRepository) Delete(ctx context.Context, id int) error {
	const op = "repository.category.delete"
	log := r.log.With(zap.String("op", op))

	query := `delete from category where id = $1`

	_, err := r.db.
		ExecContext(
			ctx,
			query,
			id,
		)

	if err != nil {
		return libs.QueryError(log, op, err)
	}

	return nil
}
