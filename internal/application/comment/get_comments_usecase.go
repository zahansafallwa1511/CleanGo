package comment

type GetCommentsUseCase struct {
	// TODO: Add repository dependency
}

func NewGetCommentsUseCase() *GetCommentsUseCase {
	return &GetCommentsUseCase{}
}

func (uc *GetCommentsUseCase) Execute(postID string) ([]CommentResponse, error) {
	// TODO: Call repository
	return []CommentResponse{
		{ID: 1, PostID: 1, UserID: 1, Content: "Great post!", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
		{ID: 2, PostID: 1, UserID: 2, Content: "Thanks for sharing", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
	}, nil
}
