package post

import (
	"context"

	"cleanandclean/internal/core/domain"
	"cleanandclean/internal/core/provider"
)

type GetPostInput struct {
	ID uint64
}

type GetPostOutput struct {
	Post *domain.Post
}

type GetPostUseCase struct{}

func NewGetPostUseCase() *GetPostUseCase {
	return &GetPostUseCase{}
}

func (uc *GetPostUseCase) Execute(ctx context.Context, input GetPostInput) (*GetPostOutput, error) {
	post, err := uc.repo().FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &GetPostOutput{Post: post}, nil
}

func (uc *GetPostUseCase) repo() PostRepository {
	return provider.Instance().GetServiceContainer().Get("PostRepository").(PostRepository)
}
