package post

import (
	"context"

	"cleanandclean/internal/core/domain"
)

type GetPostInput struct {
	ID uint64
}

type GetPostOutput struct {
	Post *domain.Post
}

type GetPostUseCase struct {
	repo PostRepository
}

func NewGetPostUseCase(repo PostRepository) *GetPostUseCase {
	return &GetPostUseCase{repo: repo}
}

func (uc *GetPostUseCase) Execute(ctx context.Context, input GetPostInput) (*GetPostOutput, error) {
	post, err := uc.repo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &GetPostOutput{Post: post}, nil
}
