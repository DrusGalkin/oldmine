package category

import (
	"context"
	"errors"
	"forum/internal/domain/models"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *CRepository) Create(ctx context.Context, category models.Category) error {
	const op = "repository.category.create"
	log := r.log.With(zap.String("op", op))

	query := `insert into category (title) values ($1) returning id`

	var id int
	if err := r.db.
		QueryRowContext(
			ctx,
			query,
			category.Title,
		).
		Scan(&id); err != nil {
		return libs.QueryError(log, op, err)
	}

	if id == 0 {
		return libs.QueryError(log, op, errors.New("Категория не создана!"))
	}

	return nil
}
