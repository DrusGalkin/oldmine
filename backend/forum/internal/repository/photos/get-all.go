package photos

import (
	"context"
	"database/sql"
	"errors"
	"forum/internal/domain/models"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *PRepository) GetAll(ctx context.Context) ([]models.Photo, error) {
	const op = "repository.photos.get-all"
	log := r.log.With(zap.String("op", op))

	query := `SELECT id, url, index, user_id FROM photos`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, libs.QueryError(log, op, err)
	}
	defer rows.Close()

	var photos []models.Photo
	for rows.Next() {
		var photo models.Photo
		if err := rows.Scan(
			&photo.ID,
			&photo.Url,
			&photo.Index,
			&photo.UserID,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}

			return nil, libs.QueryError(log, op, err)
		}
		photos = append(photos, photo)
	}

	return photos, nil
}
