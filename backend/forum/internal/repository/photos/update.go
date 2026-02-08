package photos

import (
	"context"
	"errors"
	"forum/internal/domain/model"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
)

func (r *PRepository) Update(ctx context.Context, id int, photo model.Photo) error {
	const op = "repository.photos.update"
	log := r.log.With(zap.String("op", op))

	var query string
	var args any
	if photo.Url == "" {
		query = `update photo set index = $1 where id = $2 returning id`
		args = photo.Index
	} else {
		query = `update photo set url = $1 where id = $2 returning id`
		args = photo.Url
	}

	var newId int
	if err := r.db.
		QueryRowContext(
			ctx,
			query,
			args,
		).
		Scan(&newId); err != nil {
		return libs.QueryError(log, op, err)
	}

	if newId == 0 {
		return libs.QueryError(log, op, errors.New("Данные о фотографии не изменились"))
	}

	return nil
}
