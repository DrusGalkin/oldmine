package category

import (
	"context"
	"forum/internal/domain/model"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *CRepository) Get(ctx context.Context, id int) (model.Category, error) {
	const op = "repository.category.get"
	log := r.log.With(zap.String("op", op))

	query := `SELECT id, title FROM category WHERE id = $1`

	row := r.db.
		QueryRowContext(
			ctx,
			query,
			id,
		)

	var category model.Category
	if err := row.
		Scan(
			&category.ID,
			&category.Title,
		); err != nil {
		return model.Category{}, libs.QueryError(log, op, err)
	}

	return category, nil
}
