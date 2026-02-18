package template

import (
	"context"
	"database/sql"
	"forum/internal/domain/models"
	"forum/internal/repository/photos"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

type TemplateRepository interface {
	Get(ctx context.Context, id int) (models.Template, error)
	GetAll(ctx context.Context) ([]models.Template, error)
	Create(ctx context.Context, tmp models.Template) error
	Update(ctx context.Context, id int, tmp models.Template) error
	Delete(ctx context.Context, id int) error
}

type TRepository struct {
	db    *sql.DB
	store photos.PhotoRepository
	log   *zap.Logger
}

func New(db *sql.DB, log *zap.Logger) TemplateRepository {
	return &TRepository{
		db:  db,
		log: log,
	}
}

func (t *TRepository) Get(ctx context.Context, id int) (models.Template, error) {
	const op = "repository.template.get"
	log := t.log.With(zap.String("op", op))

	query := `select id, title, html from templates where id = ?`

	row := t.db.QueryRowContext(ctx, query, id)

	var template models.Template
	if err := row.Scan(
		&template.ID,
		&template.Title,
		&template.HTML,
	); err != nil {
		return models.Template{}, libs.QueryError(log, op, err)
	}

}

func (t *TRepository) GetAll(ctx context.Context) ([]models.Template, error) {

}

func (t *TRepository) Create(ctx context.Context, tmp models.Template) error {

}

func (t *TRepository) Update(ctx context.Context, id int, tmp models.Template) error {

}

func (t *TRepository) Delete(ctx context.Context, id int) error {

}
