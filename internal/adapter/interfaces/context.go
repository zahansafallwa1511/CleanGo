package interfaces

import "net/http"

type IContext interface {
	// Request
	Request() *http.Request
	Param(key string) string
	Query(key string) string
	Header(key string) string
	Body() []byte

	// Response
	JSON(status int, data interface{})
	String(status int, data string)
	Status(status int)
	SetHeader(key, value string)

	// Flow control
	Next()
	Abort()
	AbortWithStatus(status int)
	AbortWithJSON(status int, data interface{})

	// Storage
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}
