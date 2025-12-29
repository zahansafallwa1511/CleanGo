package post

import (
	"context"

	"cleanandclean/internal/core/domain"
	"cleanandclean/internal/core/provider"
)

type UpdatePostInput struct {
	ID      uint64
	Title   string
	Content string
}

type UpdatePostOutput struct {
	Post *domain.Post
}

type UpdatePostUseCase struct{}

func NewUpdatePostUseCase() *UpdatePostUseCase {
	return &UpdatePostUseCase{}
}

func (uc *UpdatePostUseCase) Execute(ctx context.Context, input UpdatePostInput) (*UpdatePostOutput, error) {
	repo := uc.repo()

	post, err := repo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	post.Title = input.Title
	post.Content = input.Content

	if err := repo.Update(ctx, post); err != nil {
		return nil, err
	}

	return &UpdatePostOutput{Post: post}, nil
}

func (uc *UpdatePostUseCase) repo() PostRepository {
	return provider.Instance().GetServiceContainer().Get("PostRepository").(PostRepository)
}
