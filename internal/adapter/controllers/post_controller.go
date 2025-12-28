package controllers

import (
	"net/http"
	"strconv"

	"cleanandclean/internal/core/usecases/post"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	createUC *post.CreatePostUseCase
	getUC    *post.GetPostUseCase
	listUC   *post.ListPostsUseCase
	updateUC *post.UpdatePostUseCase
	deleteUC *post.DeletePostUseCase
}

func NewPostController(
	createUC *post.CreatePostUseCase,
	getUC *post.GetPostUseCase,
	listUC *post.ListPostsUseCase,
	updateUC *post.UpdatePostUseCase,
	deleteUC *post.DeletePostUseCase,
) *PostController {
	return &PostController{
		createUC: createUC,
		getUC:    getUC,
		listUC:   listUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}

type CreatePostRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	AuthorID uint64 `json:"author_id" binding:"required"`
}

func (c *PostController) Create(ctx *gin.Context) {
	var req CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := c.createUC.Execute(ctx.Request.Context(), post.CreatePostInput{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: req.AuthorID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, output.Post)
}

func (c *PostController) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	output, err := c.getUC.Execute(ctx.Request.Context(), post.GetPostInput{ID: id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output.Post)
}

func (c *PostController) List(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	output, err := c.listUC.Execute(ctx.Request.Context(), post.ListPostsInput{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output.Posts)
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (c *PostController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req UpdatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := c.updateUC.Execute(ctx.Request.Context(), post.UpdatePostInput{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output.Post)
}

func (c *PostController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := c.deleteUC.Execute(ctx.Request.Context(), post.DeletePostInput{ID: id}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
