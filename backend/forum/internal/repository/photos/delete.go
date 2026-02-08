package photos

import (
	"context"
	"github.com/DrusGalkin/libs"
	"go.uber.org/zap"
	"os"
)

func (r *PRepository) Delete(ctx context.Context, id int) error {
	const op = "repository.photos.delete"
	log := r.log.With(zap.String("op", op))

	var query = `delete from photos where id = $1`

	photo, err := r.Get(ctx, id)
	if err != nil {
		return libs.QueryError(log, op, err)
	}

	_, err = r.db.ExecContext(ctx, query, photo.ID)
	if err != nil {
		return libs.QueryError(log, op, err)
	}

	os.Remove(photo.Url)
	return nil
}
