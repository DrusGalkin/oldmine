package photos

import (
	"context"
	"errors"
	"forum/internal/domain/models"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *PRepository) Create(ctx context.Context, photo models.Photo) error {
	const op = "repository.photos.create"
	log := r.log.With(zap.String("op", op))

	query := `insert into photos (url, index, user_id) values  ($1, $2, $3) returning id`

	var id int
	if err := r.db.
		QueryRowContext(
			ctx,
			query,
			photo.Url,
			photo.Index,
			photo.UserID,
		).
		Scan(&id); err != nil {
		return libs.QueryError(log, op, err)
	}

	if id == 0 {
		return libs.QueryError(log, op, errors.New("Записть о фотографии не сохранена в базе данных"))
	}

	return nil
}
