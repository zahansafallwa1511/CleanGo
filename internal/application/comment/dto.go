package comment

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1"`
	UserID  int    `json:"user_id" binding:"required,gt=0"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1"`
}

type CommentResponse struct {
	ID        int    `json:"id"`
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
