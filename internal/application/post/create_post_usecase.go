package post

type CreatePostUseCase struct {
	// TODO: Add repository dependency
	// repo IPostRepository
}

func NewCreatePostUseCase() *CreatePostUseCase {
	return &CreatePostUseCase{}
}

func (uc *CreatePostUseCase) Execute(req CreatePostRequest) (*PostResponse, error) {
	// TODO: Call repository
	return &PostResponse{
		ID:        1,
		Title:     req.Title,
		Content:   req.Content,
		UserID:    req.UserID,
		CreatedAt: "2025-01-01T00:00:00Z",
		UpdatedAt: "2025-01-01T00:00:00Z",
	}, nil
}
