package interfaces

import "context"

type IRouter interface {
	Run(addr string) error
	Shutdown(ctx context.Context) error
	Close() error
	GET(path string, handler interface{})
	POST(path string, handler interface{})
	PUT(path string, handler interface{})
	DELETE(path string, handler interface{})
	PATCH(path string, handler interface{})
	Group(prefix string) IRouter
}
