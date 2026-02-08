package category

import (
	"context"
	"errors"
	"forum/internal/domain/model"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *CRepository) Update(ctx context.Context, category model.Category) error {
	const op = "repository.category.update"
	log := r.log.With(zap.String("op", op))

	query := `update category set title = $1 where id = $2 returning id`

	var id int
	if err := r.db.
		QueryRowContext(
			ctx,
			query,
			category.Title,
			category.ID,
		).
		Scan(&id); err != nil {
		return libs.QueryError(log, op, err)
	}

	if id == 0 {
		return libs.QueryError(log, op, errors.New("Категория не обновлена!"))
	}

	return nil
}
