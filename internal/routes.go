package internal

import (
	"cleanandclean/internal/adapter/interfaces"
	"cleanandclean/internal/application/comment"
	"cleanandclean/internal/application/health"
	"cleanandclean/internal/application/post"
)

type Route struct {
	Method   string
	Path     string
	Handlers []interfaces.HandlerFunc
}

func GetRoutes() []Route {
	return []Route{
		// Health
		{"GET", "/health", []interfaces.HandlerFunc{health.Readiness}},

		// Posts
		{"GET", "/posts", []interfaces.HandlerFunc{post.GetPosts}},
		{"GET", "/posts/:id", []interfaces.HandlerFunc{post.GetPost}},
		{"POST", "/posts", []interfaces.HandlerFunc{post.CreatePost}},
		{"PUT", "/posts/:id", []interfaces.HandlerFunc{post.UpdatePost}},
		{"DELETE", "/posts/:id", []interfaces.HandlerFunc{post.DeletePost}},

		// Comments
		{"GET", "/posts/:id/comments", []interfaces.HandlerFunc{comment.GetComments}},
		{"POST", "/posts/:id/comments", []interfaces.HandlerFunc{comment.CreateComment}},
		{"PUT", "/comments/:id", []interfaces.HandlerFunc{comment.UpdateComment}},
		{"DELETE", "/comments/:id", []interfaces.HandlerFunc{comment.DeleteComment}},
	}
}
