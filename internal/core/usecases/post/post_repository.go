package post

import (
	"context"

	"cleanandclean/internal/core/domain"
)

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	FindByID(ctx context.Context, id uint64) (*domain.Post, error)
	FindAll(ctx context.Context, limit, offset int) ([]*domain.Post, error)
	Update(ctx context.Context, post *domain.Post) error
	Delete(ctx context.Context, id uint64) error
}
