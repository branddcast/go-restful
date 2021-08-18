package article

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Article, error)
	GetOne(ctx context.Context, id uint) (Article, error)
	Create(ctx context.Context, article *Article) error
	Update(ctx context.Context, id uint, article Article) error
	Delete(ctx context.Context, id uint) error
}
