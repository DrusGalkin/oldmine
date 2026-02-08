package photos

import (
	"context"
	"forum/internal/domain/model"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *PRepository) Get(ctx context.Context, id int) (model.Photo, error) {
	const op = "repository.photos.get-all"
	log := r.log.With(zap.String("op", op))

	query := `SELECT id, url, index, user_id FROM photos WHERE id = $1`

	var photo model.Photo
	if err := r.db.
		QueryRowContext(
			ctx,
			query,
			id,
		).Scan(
		&photo.ID,
		&photo.Url,
		&photo.Index,
		&photo.UserID,
	); err != nil {
		return model.Photo{}, libs.QueryError(log, op, err)
	}

	return photo, nil
}
