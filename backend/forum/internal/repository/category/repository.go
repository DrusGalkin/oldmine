package category

import (
	"context"
	"forum/internal/domain/model"
)

type CategoryRepository interface {
	Get(ctx context.Context, id int) (model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)
	GetByTemplateId(ctx context.Context, templateId int) ([]model.Category, error)
	Create(ctx context.Context, category model.Category) error
	Update(ctx context.Context, id int, category model.Category) error
	Delete(ctx context.Context, id int) error
}
