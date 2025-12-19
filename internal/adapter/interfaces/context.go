package interfaces

import "net/http"

// IContext defines the contract for HTTP request/response handling
type IContext interface {
	// Request
	Request() *http.Request
	Param(key string) string
	Query(key string) string
	QueryInt(key string, defaultVal int) int
	Header(key string) string
	Body() []byte

	// Request binding
	Bind(obj interface{}) error
	BindJSON(obj interface{}) error
	BindQuery(obj interface{}) error

	// Response - raw
	JSON(status int, data interface{})
	String(status int, data string)
	Status(status int)
	SetHeader(key, value string)

	// Response - structured
	Success(status int, data interface{})
	SuccessWithMeta(status int, data interface{}, page, perPage, total, totalPages int)
	Error(status int, code string, message string)
	ErrorWithDetails(status int, code string, message string, details map[string]string)
	ValidationError(details map[string]string)

	// Flow control
	Next()
	Abort()
	AbortWithStatus(status int)
	AbortWithJSON(status int, data interface{})
	AbortWithError(status int, code string, message string)

	// Storage
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

// HandlerFunc is the signature for HTTP handlers
type HandlerFunc func(ctx IContext)
