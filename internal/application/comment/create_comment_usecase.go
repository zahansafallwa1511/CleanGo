package comment

type CreateCommentUseCase struct {
	// TODO: Add repository dependency
}

func NewCreateCommentUseCase() *CreateCommentUseCase {
	return &CreateCommentUseCase{}
}

func (uc *CreateCommentUseCase) Execute(postID string, req CreateCommentRequest) (*CommentResponse, error) {
	// TODO: Call repository
	return &CommentResponse{
		ID:        1,
		PostID:    1,
		UserID:    req.UserID,
		Content:   req.Content,
		CreatedAt: "2025-01-01T00:00:00Z",
		UpdatedAt: "2025-01-01T00:00:00Z",
	}, nil
}
