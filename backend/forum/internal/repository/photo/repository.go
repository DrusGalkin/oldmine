package photo

import (
	"context"
	"forum/internal/domain/model"
)

type PhotoRepository interface {
	Get(ctx context.Context, id int) (model.Photo, error)
	GetAll(ctx context.Context, templateID int) ([]model.Photo, error)
	Create(ctx context.Context, photo model.Photo) error
	Update(ctx context.Context, id int, photo model.Photo) error
	Delete(ctx context.Context, id int) error
}
