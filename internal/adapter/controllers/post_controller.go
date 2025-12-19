package controllers

import (
	"errors"
	"net/http"

	httpresponse "cleanandclean/internal/adapter/http"
	"cleanandclean/internal/adapter/interfaces"
	"cleanandclean/internal/application/post"
)

type PostController struct {
	getPostsUseCase  *post.GetPostsUseCase
	getPostUseCase   *post.GetPostUseCase
	createPostUseCase *post.CreatePostUseCase
	updatePostUseCase *post.UpdatePostUseCase
	deletePostUseCase *post.DeletePostUseCase
}

func NewPostController(
	getPostsUseCase *post.GetPostsUseCase,
	getPostUseCase *post.GetPostUseCase,
	createPostUseCase *post.CreatePostUseCase,
	updatePostUseCase *post.UpdatePostUseCase,
	deletePostUseCase *post.DeletePostUseCase,
) *PostController {
	return &PostController{
		getPostsUseCase:   getPostsUseCase,
		getPostUseCase:    getPostUseCase,
		createPostUseCase: createPostUseCase,
		updatePostUseCase: updatePostUseCase,
		deletePostUseCase: deletePostUseCase,
	}
}

func (c *PostController) GetAll(ctx interfaces.IContext) {
	page := ctx.QueryInt("page", 1)
	perPage := ctx.QueryInt("per_page", 10)

	posts, total, err := c.getPostsUseCase.Execute(page, perPage)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, httpresponse.ErrCodeInternal, "Failed to fetch posts")
		return
	}

	totalPages := (total + perPage - 1) / perPage
	ctx.SuccessWithMeta(http.StatusOK, posts, page, perPage, total, totalPages)
}

func (c *PostController) GetByID(ctx interfaces.IContext) {
	id := ctx.Param("id")

	result, err := c.getPostUseCase.Execute(id)
	if err != nil {
		if errors.Is(err, post.ErrPostNotFound) {
			ctx.Error(http.StatusNotFound, httpresponse.ErrCodeNotFound, "Post not found")
			return
		}
		ctx.Error(http.StatusInternalServerError, httpresponse.ErrCodeInternal, "Failed to fetch post")
		return
	}

	ctx.Success(http.StatusOK, result)
}

func (c *PostController) Create(ctx interfaces.IContext) {
	var req post.CreatePostRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.ValidationError(map[string]string{"error": err.Error()})
		return
	}

	result, err := c.createPostUseCase.Execute(req)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, httpresponse.ErrCodeInternal, "Failed to create post")
		return
	}

	ctx.Success(http.StatusCreated, result)
}

func (c *PostController) Update(ctx interfaces.IContext) {
	id := ctx.Param("id")

	var req post.UpdatePostRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.ValidationError(map[string]string{"error": err.Error()})
		return
	}

	result, err := c.updatePostUseCase.Execute(id, req)
	if err != nil {
		if errors.Is(err, post.ErrPostNotFound) {
			ctx.Error(http.StatusNotFound, httpresponse.ErrCodeNotFound, "Post not found")
			return
		}
		ctx.Error(http.StatusInternalServerError, httpresponse.ErrCodeInternal, "Failed to update post")
		return
	}

	ctx.Success(http.StatusOK, result)
}

func (c *PostController) Delete(ctx interfaces.IContext) {
	id := ctx.Param("id")

	err := c.deletePostUseCase.Execute(id)
	if err != nil {
		if errors.Is(err, post.ErrPostNotFound) {
			ctx.Error(http.StatusNotFound, httpresponse.ErrCodeNotFound, "Post not found")
			return
		}
		ctx.Error(http.StatusInternalServerError, httpresponse.ErrCodeInternal, "Failed to delete post")
		return
	}

	ctx.Success(http.StatusNoContent, nil)
}
