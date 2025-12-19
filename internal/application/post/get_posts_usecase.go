package post

type GetPostsUseCase struct {
	// TODO: Add repository dependency
	// repo IPostRepository
}

func NewGetPostsUseCase() *GetPostsUseCase {
	return &GetPostsUseCase{}
}

func (uc *GetPostsUseCase) Execute(page, perPage int) ([]PostResponse, int, error) {
	// TODO: Call repository
	posts := []PostResponse{
		{ID: 1, Title: "First Post", Content: "Content", UserID: 1, CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
		{ID: 2, Title: "Second Post", Content: "Content", UserID: 1, CreatedAt: "2025-01-02T00:00:00Z", UpdatedAt: "2025-01-02T00:00:00Z"},
	}
	total := 25
	return posts, total, nil
}
