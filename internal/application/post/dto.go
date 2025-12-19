package post

// CreatePostRequest represents the request body for creating a post
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=255"`
	Content string `json:"content" binding:"required,min=1"`
	UserID  int    `json:"user_id" binding:"required,gt=0"`
}

// UpdatePostRequest represents the request body for updating a post
type UpdatePostRequest struct {
	Title   string `json:"title" binding:"omitempty,min=1,max=255"`
	Content string `json:"content" binding:"omitempty,min=1"`
}

// PostResponse represents a post in API responses
type PostResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ListPostsQuery represents query parameters for listing posts
type ListPostsQuery struct {
	Page    int `form:"page" binding:"omitempty,min=1"`
	PerPage int `form:"per_page" binding:"omitempty,min=1,max=100"`
}
