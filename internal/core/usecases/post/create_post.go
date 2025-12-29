package post

import (
	"context"

	"cleanandclean/internal/core/domain"
	"cleanandclean/internal/core/provider"
)

type CreatePostInput struct {
	Title    string
	Content  string
	AuthorID uint64
}

type CreatePostOutput struct {
	Post *domain.Post
}

type CreatePostUseCase struct{}

func NewCreatePostUseCase() *CreatePostUseCase {
	return &CreatePostUseCase{}
}

func (uc *CreatePostUseCase) Execute(ctx context.Context, input CreatePostInput) (*CreatePostOutput, error) {
	repo := uc.repo()

	post := &domain.Post{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: input.AuthorID,
	}

	if err := repo.Create(ctx, post); err != nil {
		return nil, err
	}

	return &CreatePostOutput{Post: post}, nil
}

func (uc *CreatePostUseCase) repo() PostRepository {
	return provider.Instance().GetServiceContainer().Get("PostRepository").(PostRepository)
}
