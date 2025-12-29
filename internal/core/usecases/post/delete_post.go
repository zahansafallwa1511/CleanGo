package post

import (
	"context"

	"cleanandclean/internal/core/provider"
)

type DeletePostInput struct {
	ID uint64
}

type DeletePostUseCase struct{}

func NewDeletePostUseCase() *DeletePostUseCase {
	return &DeletePostUseCase{}
}

func (uc *DeletePostUseCase) Execute(ctx context.Context, input DeletePostInput) error {
	return uc.repo().Delete(ctx, input.ID)
}

func (uc *DeletePostUseCase) repo() PostRepository {
	return provider.Instance().GetServiceContainer().Get("PostRepository").(PostRepository)
}
