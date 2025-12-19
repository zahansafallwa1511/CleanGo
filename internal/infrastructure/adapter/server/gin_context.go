package server

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	httpresponse "cleanandclean/internal/adapter/http"
)

// GinContext implements IContext using Gin's context
type GinContext struct {
	c *gin.Context
}

// Request methods
func (ctx *GinContext) Request() *http.Request  { return ctx.c.Request }
func (ctx *GinContext) Param(key string) string { return ctx.c.Param(key) }
func (ctx *GinContext) Query(key string) string { return ctx.c.Query(key) }
func (ctx *GinContext) QueryInt(key string, defaultVal int) int {
	val := ctx.c.Query(key)
	if val == "" {
		return defaultVal
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return intVal
}
func (ctx *GinContext) Header(key string) string { return ctx.c.GetHeader(key) }
func (ctx *GinContext) Body() []byte             { b, _ := io.ReadAll(ctx.c.Request.Body); return b }

// Request binding
func (ctx *GinContext) Bind(obj interface{}) error      { return ctx.c.ShouldBind(obj) }
func (ctx *GinContext) BindJSON(obj interface{}) error  { return ctx.c.ShouldBindJSON(obj) }
func (ctx *GinContext) BindQuery(obj interface{}) error { return ctx.c.ShouldBindQuery(obj) }

// Response - raw
func (ctx *GinContext) JSON(status int, data interface{}) { ctx.c.JSON(status, data) }
func (ctx *GinContext) String(status int, data string)    { ctx.c.String(status, data) }
func (ctx *GinContext) Status(status int)                 { ctx.c.Status(status) }
func (ctx *GinContext) SetHeader(key, value string)       { ctx.c.Header(key, value) }

// Response - structured
func (ctx *GinContext) Success(status int, data interface{}) {
	ctx.c.JSON(status, httpresponse.APIResponse{
		Success: true,
		Data:    data,
	})
}

func (ctx *GinContext) SuccessWithMeta(status int, data interface{}, page, perPage, total, totalPages int) {
	ctx.c.JSON(status, httpresponse.APIResponse{
		Success: true,
		Data:    data,
		Meta: &httpresponse.APIMeta{
			Page:       page,
			PerPage:    perPage,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

func (ctx *GinContext) Error(status int, code string, message string) {
	ctx.c.JSON(status, httpresponse.APIResponse{
		Success: false,
		Error: &httpresponse.APIError{
			Code:    code,
			Message: message,
		},
	})
}

func (ctx *GinContext) ErrorWithDetails(status int, code string, message string, details map[string]string) {
	ctx.c.JSON(status, httpresponse.APIResponse{
		Success: false,
		Error: &httpresponse.APIError{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

func (ctx *GinContext) ValidationError(details map[string]string) {
	ctx.c.JSON(http.StatusBadRequest, httpresponse.APIResponse{
		Success: false,
		Error: &httpresponse.APIError{
			Code:    httpresponse.ErrCodeValidation,
			Message: "Validation failed",
			Details: details,
		},
	})
}

// Flow control
func (ctx *GinContext) Next()                      { ctx.c.Next() }
func (ctx *GinContext) Abort()                     { ctx.c.Abort() }
func (ctx *GinContext) AbortWithStatus(status int) { ctx.c.AbortWithStatus(status) }
func (ctx *GinContext) AbortWithJSON(status int, data interface{}) {
	ctx.c.AbortWithStatusJSON(status, data)
}
func (ctx *GinContext) AbortWithError(status int, code string, message string) {
	ctx.c.AbortWithStatusJSON(status, httpresponse.APIResponse{
		Success: false,
		Error: &httpresponse.APIError{
			Code:    code,
			Message: message,
		},
	})
}

// Storage
func (ctx *GinContext) Set(key string, value interface{})  { ctx.c.Set(key, value) }
func (ctx *GinContext) Get(key string) (interface{}, bool) { return ctx.c.Get(key) }
