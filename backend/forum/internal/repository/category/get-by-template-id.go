package category

import (
	"context"
	"database/sql"
	"errors"
	"forum/internal/domain/model"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *CRepository) GetByTemplateId(ctx context.Context, templateId int) ([]model.Category, error) {
	const op = "repository.category.get-by-template-id"
	log := r.log.With(zap.String("op", op))

	query := `SELECT id, title FROM category where template_id = $1`

	rows, err := r.db.
		QueryContext(
			ctx,
			query,
			templateId,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, libs.QueryError(log, op, err)
	}
	defer rows.Close()

	var categorys []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.
			Scan(
				&category.ID,
				&category.Title,
			); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}
			return nil, libs.QueryError(log, op, err)
		}
		categorys = append(categorys, category)
	}

	return categorys, nil
}
