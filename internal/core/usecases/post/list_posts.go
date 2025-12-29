package post

import (
	"context"

	"cleanandclean/internal/core/domain"
	"cleanandclean/internal/core/provider"
)

type ListPostsInput struct {
	Limit  int
	Offset int
}

type ListPostsOutput struct {
	Posts []*domain.Post
}

type ListPostsUseCase struct{}

func NewListPostsUseCase() *ListPostsUseCase {
	return &ListPostsUseCase{}
}

func (uc *ListPostsUseCase) Execute(ctx context.Context, input ListPostsInput) (*ListPostsOutput, error) {
	posts, err := uc.repo().FindAll(ctx, input.Limit, input.Offset)
	if err != nil {
		return nil, err
	}

	return &ListPostsOutput{Posts: posts}, nil
}

func (uc *ListPostsUseCase) repo() PostRepository {
	return provider.Instance().GetServiceContainer().Get("PostRepository").(PostRepository)
}
