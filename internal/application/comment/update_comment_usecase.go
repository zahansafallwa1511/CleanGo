package comment

import "errors"

var ErrCommentNotFound = errors.New("comment not found")

type UpdateCommentUseCase struct {
	// TODO: Add repository dependency
}

func NewUpdateCommentUseCase() *UpdateCommentUseCase {
	return &UpdateCommentUseCase{}
}

func (uc *UpdateCommentUseCase) Execute(id string, req UpdateCommentRequest) (*CommentResponse, error) {
	if id == "0" {
		return nil, ErrCommentNotFound
	}
	// TODO: Call repository
	return &CommentResponse{
		ID:        1,
		PostID:    1,
		UserID:    1,
		Content:   req.Content,
		CreatedAt: "2025-01-01T00:00:00Z",
		UpdatedAt: "2025-01-01T00:00:00Z",
	}, nil
}
