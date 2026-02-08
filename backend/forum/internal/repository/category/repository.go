package category

import (
	"context"
	"database/sql"
	"errors"
	"forum/internal/domain/model"
	"go.uber.org/zap"
)

type CategoryRepository interface {
	Get(ctx context.Context, id int) (model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)
	GetByTemplateId(ctx context.Context, templateId int) ([]model.Category, error)
	Create(ctx context.Context, category model.Category) error
	Update(ctx context.Context, id int, category model.Category) error
	Delete(ctx context.Context, id int) error
}

type CRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func New(db *sql.DB, log *zap.Logger) CategoryRepository {
	return &CRepository{
		db:  db,
		log: log,
	}
}

func (r *CRepository) Get(ctx context.Context, id int) (model.Category, error) {
	const op = "repository.category.get"
	log := r.log.With(zap.String("op", op))

	query := `SELECT id, title FROM category WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)

	var category model.Category
	if err := row.Scan(
		&category.ID,
		&category.Title,
	); err != nil {

	}

	return category, nil
}

func (r *CRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	const op = "repository.category.get-all"
	log := r.log.With(zap.String("op", op))

	query := `SELECT id, title FROM category`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	var categorys []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(
			&category.ID,
			&category.Title,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}
			return nil, err
		}
		categorys = append(categorys, category)
	}

	return categorys, nil
}

func (r *CRepository) GetByTemplateId(ctx context.Context, templateId int) ([]model.Category, error) {

}

func (r *CRepository) Create(ctx context.Context, category model.Category) error {

}

func (r *CRepository) Update(ctx context.Context, id int, category model.Category) error {

}

func (r *CRepository) Delete(ctx context.Context, id int) error {

}
