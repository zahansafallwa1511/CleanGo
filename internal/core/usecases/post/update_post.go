package post

import (
	"context"

	"cleanandclean/internal/core/domain"
)

type UpdatePostInput struct {
	ID      uint64
	Title   string
	Content string
}

type UpdatePostOutput struct {
	Post *domain.Post
}

type UpdatePostUseCase struct {
	repo PostRepository
}

func NewUpdatePostUseCase(repo PostRepository) *UpdatePostUseCase {
	return &UpdatePostUseCase{repo: repo}
}

func (uc *UpdatePostUseCase) Execute(ctx context.Context, input UpdatePostInput) (*UpdatePostOutput, error) {
	post, err := uc.repo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	post.Title = input.Title
	post.Content = input.Content

	if err := uc.repo.Update(ctx, post); err != nil {
		return nil, err
	}

	return &UpdatePostOutput{Post: post}, nil
}
