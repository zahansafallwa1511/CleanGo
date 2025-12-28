package post

import (
	"context"

	"cleanandclean/internal/core/domain"
)

type ListPostsInput struct {
	Limit  int
	Offset int
}

type ListPostsOutput struct {
	Posts []*domain.Post
}

type ListPostsUseCase struct {
	repo PostRepository
}

func NewListPostsUseCase(repo PostRepository) *ListPostsUseCase {
	return &ListPostsUseCase{repo: repo}
}

func (uc *ListPostsUseCase) Execute(ctx context.Context, input ListPostsInput) (*ListPostsOutput, error) {
	posts, err := uc.repo.FindAll(ctx, input.Limit, input.Offset)
	if err != nil {
		return nil, err
	}

	return &ListPostsOutput{Posts: posts}, nil
}
