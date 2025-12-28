package post

import "context"

type DeletePostInput struct {
	ID uint64
}

type DeletePostUseCase struct {
	repo PostRepository
}

func NewDeletePostUseCase(repo PostRepository) *DeletePostUseCase {
	return &DeletePostUseCase{repo: repo}
}

func (uc *DeletePostUseCase) Execute(ctx context.Context, input DeletePostInput) error {
	return uc.repo.Delete(ctx, input.ID)
}
