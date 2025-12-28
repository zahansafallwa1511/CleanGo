package post

import (
	"context"

	"cleanandclean/internal/core/domain"
)

type CreatePostInput struct {
	Title    string
	Content  string
	AuthorID uint64
}

type CreatePostOutput struct {
	Post *domain.Post
}

type CreatePostUseCase struct {
	repo PostRepository
}

func NewCreatePostUseCase(repo PostRepository) *CreatePostUseCase {
	return &CreatePostUseCase{repo: repo}
}

func (uc *CreatePostUseCase) Execute(ctx context.Context, input CreatePostInput) (*CreatePostOutput, error) {
	post := &domain.Post{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: input.AuthorID,
	}

	if err := uc.repo.Create(ctx, post); err != nil {
		return nil, err
	}

	return &CreatePostOutput{Post: post}, nil
}
