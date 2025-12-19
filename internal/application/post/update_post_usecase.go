package post

type UpdatePostUseCase struct {
	// TODO: Add repository dependency
	// repo IPostRepository
}

func NewUpdatePostUseCase() *UpdatePostUseCase {
	return &UpdatePostUseCase{}
}

func (uc *UpdatePostUseCase) Execute(id string, req UpdatePostRequest) (*PostResponse, error) {
	if id == "0" {
		return nil, ErrPostNotFound
	}
	// TODO: Call repository
	return &PostResponse{
		ID:        1,
		Title:     req.Title,
		Content:   req.Content,
		UserID:    1,
		CreatedAt: "2025-01-01T00:00:00Z",
		UpdatedAt: "2025-01-01T00:00:00Z",
	}, nil
}
