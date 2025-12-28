package adapters

import (
	"context"
	"net/http"

	"cleanandclean/internal/adapter/interfaces"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	engine *gin.Engine
	group  *gin.RouterGroup
	server *http.Server
}

func NewGinRouter() *GinRouter {
	engine := gin.Default()
	return &GinRouter{
		engine: engine,
		group:  engine.Group(""),
	}
}

func (r *GinRouter) Run(addr string) error {
	r.server = &http.Server{
		Addr:    addr,
		Handler: r.engine,
	}
	return r.server.ListenAndServe()
}

func (r *GinRouter) Shutdown(ctx context.Context) error {
	if r.server != nil {
		return r.server.Shutdown(ctx)
	}
	return nil
}

func (r *GinRouter) Close() error {
	if r.server != nil {
		return r.server.Close()
	}
	return nil
}

func (r *GinRouter) GET(path string, handler interface{}) {
	r.group.GET(path, toGinHandler(handler))
}

func (r *GinRouter) POST(path string, handler interface{}) {
	r.group.POST(path, toGinHandler(handler))
}

func (r *GinRouter) PUT(path string, handler interface{}) {
	r.group.PUT(path, toGinHandler(handler))
}

func (r *GinRouter) DELETE(path string, handler interface{}) {
	r.group.DELETE(path, toGinHandler(handler))
}

func (r *GinRouter) PATCH(path string, handler interface{}) {
	r.group.PATCH(path, toGinHandler(handler))
}

func toGinHandler(handler interface{}) gin.HandlerFunc {
	switch h := handler.(type) {
	case gin.HandlerFunc:
		return h
	case func(*gin.Context):
		return h
	default:
		panic("unsupported handler type")
	}
}

func (r *GinRouter) Group(prefix string) interfaces.IRouter {
	return &GinRouter{
		engine: r.engine,
		group:  r.group.Group(prefix),
	}
}

var _ interfaces.IRouter = (*GinRouter)(nil)
