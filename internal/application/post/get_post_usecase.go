package post

import "errors"

var ErrPostNotFound = errors.New("post not found")

type GetPostUseCase struct {
	// TODO: Add repository dependency
	// repo IPostRepository
}

func NewGetPostUseCase() *GetPostUseCase {
	return &GetPostUseCase{}
}

func (uc *GetPostUseCase) Execute(id string) (*PostResponse, error) {
	if id == "0" {
		return nil, ErrPostNotFound
	}
	// TODO: Call repository
	return &PostResponse{
		ID:        1,
		Title:     "Sample Post",
		Content:   "Sample content",
		UserID:    1,
		CreatedAt: "2025-01-01T00:00:00Z",
		UpdatedAt: "2025-01-01T00:00:00Z",
	}, nil
}
