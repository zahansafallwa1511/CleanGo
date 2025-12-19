package controllers

import (
	"errors"
	"net/http"

	httpresponse "cleanandclean/internal/adapter/http"
	"cleanandclean/internal/adapter/interfaces"
	"cleanandclean/internal/application/comment"
)

type CommentController struct {
	getCommentsUseCase   *comment.GetCommentsUseCase
	createCommentUseCase *comment.CreateCommentUseCase
	updateCommentUseCase *comment.UpdateCommentUseCase
	deleteCommentUseCase *comment.DeleteCommentUseCase
}

func NewCommentController(
	getCommentsUseCase *comment.GetCommentsUseCase,
	createCommentUseCase *comment.CreateCommentUseCase,
	updateCommentUseCase *comment.UpdateCommentUseCase,
	deleteCommentUseCase *comment.DeleteCommentUseCase,
) *CommentController {
	return &CommentController{
		getCommentsUseCase:   getCommentsUseCase,
		createCommentUseCase: createCommentUseCase,
		updateCommentUseCase: updateCommentUseCase,
		deleteCommentUseCase: deleteCommentUseCase,
	}
}

func (c *CommentController) GetByPostID(ctx interfaces.IContext) {
	postID := ctx.Param("id")

	comments, err := c.getCommentsUseCase.Execute(postID)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, httpresponse.ErrCodeInternal, "Failed to fetch comments")
		return
	}

	ctx.Success(http.StatusOK, comments)
}

func (c *CommentController) Create(ctx interfaces.IContext) {
	postID := ctx.Param("id")

	var req comment.CreateCommentRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.ValidationError(map[string]string{"error": err.Error()})
		return
	}

	result, err := c.createCommentUseCase.Execute(postID, req)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, httpresponse.ErrCodeInternal, "Failed to create comment")
		return
	}

	ctx.Success(http.StatusCreated, result)
}

func (c *CommentController) Update(ctx interfaces.IContext) {
	id := ctx.Param("id")

	var req comment.UpdateCommentRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.ValidationError(map[string]string{"error": err.Error()})
		return
	}

	result, err := c.updateCommentUseCase.Execute(id, req)
	if err != nil {
		if errors.Is(err, comment.ErrCommentNotFound) {
			ctx.Error(http.StatusNotFound, httpresponse.ErrCodeNotFound, "Comment not found")
			return
		}
		ctx.Error(http.StatusInternalServerError, httpresponse.ErrCodeInternal, "Failed to update comment")
		return
	}

	ctx.Success(http.StatusOK, result)
}

func (c *CommentController) Delete(ctx interfaces.IContext) {
	id := ctx.Param("id")

	err := c.deleteCommentUseCase.Execute(id)
	if err != nil {
		if errors.Is(err, comment.ErrCommentNotFound) {
			ctx.Error(http.StatusNotFound, httpresponse.ErrCodeNotFound, "Comment not found")
			return
		}
		ctx.Error(http.StatusInternalServerError, httpresponse.ErrCodeInternal, "Failed to delete comment")
		return
	}

	ctx.Success(http.StatusNoContent, nil)
}
