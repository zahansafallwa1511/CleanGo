package internal

import (
	"cleanandclean/internal/adapter/controllers"
	"cleanandclean/internal/adapter/interfaces"
	"cleanandclean/internal/application/comment"
	"cleanandclean/internal/application/health"
	"cleanandclean/internal/application/post"
)

// RegisterRoutes sets up all application routes
func RegisterRoutes(server interfaces.IServer) {
	// Initialize post use cases
	getPostsUseCase := post.NewGetPostsUseCase()
	getPostUseCase := post.NewGetPostUseCase()
	createPostUseCase := post.NewCreatePostUseCase()
	updatePostUseCase := post.NewUpdatePostUseCase()
	deletePostUseCase := post.NewDeletePostUseCase()

	// Initialize comment use cases
	getCommentsUseCase := comment.NewGetCommentsUseCase()
	createCommentUseCase := comment.NewCreateCommentUseCase()
	updateCommentUseCase := comment.NewUpdateCommentUseCase()
	deleteCommentUseCase := comment.NewDeleteCommentUseCase()

	// Initialize controllers
	postController := controllers.NewPostController(
		getPostsUseCase,
		getPostUseCase,
		createPostUseCase,
		updatePostUseCase,
		deletePostUseCase,
	)
	commentController := controllers.NewCommentController(
		getCommentsUseCase,
		createCommentUseCase,
		updateCommentUseCase,
		deleteCommentUseCase,
	)

	// Health check
	server.RegisterRoute("GET", "/health", health.Readiness)

	// API v1
	v1 := server.Group("/api/v1")
	{
		// Posts
		posts := v1.Group("/posts")
		{
			posts.RegisterRoute("GET", "", postController.GetAll)
			posts.RegisterRoute("GET", "/:id", postController.GetByID)
			posts.RegisterRoute("POST", "", postController.Create)
			posts.RegisterRoute("PUT", "/:id", postController.Update)
			posts.RegisterRoute("DELETE", "/:id", postController.Delete)

			// Comments under posts
			posts.RegisterRoute("GET", "/:id/comments", commentController.GetByPostID)
			posts.RegisterRoute("POST", "/:id/comments", commentController.Create)
		}

		// Comments
		comments := v1.Group("/comments")
		{
			comments.RegisterRoute("PUT", "/:id", commentController.Update)
			comments.RegisterRoute("DELETE", "/:id", commentController.Delete)
		}
	}
}
