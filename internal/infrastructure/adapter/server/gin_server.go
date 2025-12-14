package server

import (
	"context"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"cleanandclean/internal/adapter/interfaces"
)

type GinServer struct {
	addr       string
	engine     *gin.Engine
	httpServer *http.Server
}

func NewGinServer(addr string) *GinServer {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	return &GinServer{
		addr:   addr,
		engine: engine,
	}
}

func (s *GinServer) Start() error {
	s.httpServer = &http.Server{
		Addr:    s.addr,
		Handler: s.engine,
	}
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Server error - logged at project level if needed
		}
	}()
	return nil
}

func (s *GinServer) Stop(ctx context.Context) error {
	if s.httpServer != nil {
		return s.httpServer.Shutdown(ctx)
	}
	return nil
}

func (s *GinServer) Addr() string {
	return s.addr
}

func (s *GinServer) Use(handlers ...interfaces.HandlerFunc) {
	s.engine.Use(toGinHandlers(handlers)...)
}

func (s *GinServer) RegisterRoute(method, path string, handlers ...interfaces.HandlerFunc) {
	s.engine.Handle(method, path, toGinHandlers(handlers)...)
}

func (s *GinServer) Group(prefix string, handlers ...interfaces.HandlerFunc) interfaces.IRouteGroup {
	return &GinRouteGroup{group: s.engine.Group(prefix, toGinHandlers(handlers)...)}
}

func (s *GinServer) EnableCORS(config interfaces.CORSConfig) {
	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     config.AllowOrigins,
		AllowMethods:     config.AllowMethods,
		AllowHeaders:     config.AllowHeaders,
		AllowCredentials: config.AllowCredentials,
	}))
}

func (s *GinServer) EnableRateLimiting(config interfaces.RateLimitConfig) {}

// GinRouteGroup
type GinRouteGroup struct {
	group *gin.RouterGroup
}

func (g *GinRouteGroup) Use(handlers ...interfaces.HandlerFunc) {
	g.group.Use(toGinHandlers(handlers)...)
}

func (g *GinRouteGroup) RegisterRoute(method, path string, handlers ...interfaces.HandlerFunc) {
	g.group.Handle(method, path, toGinHandlers(handlers)...)
}

func (g *GinRouteGroup) Group(prefix string, handlers ...interfaces.HandlerFunc) interfaces.IRouteGroup {
	return &GinRouteGroup{group: g.group.Group(prefix, toGinHandlers(handlers)...)}
}

func toGinHandlers(handlers []interfaces.HandlerFunc) []gin.HandlerFunc {
	result := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		handler := h
		result[i] = func(c *gin.Context) {
			handler(&GinContext{c: c})
		}
	}
	return result
}

// GinContext
type GinContext struct {
	c *gin.Context
}

func (ctx *GinContext) Request() *http.Request                             { return ctx.c.Request }
func (ctx *GinContext) Param(key string) string                            { return ctx.c.Param(key) }
func (ctx *GinContext) Query(key string) string                            { return ctx.c.Query(key) }
func (ctx *GinContext) Header(key string) string                           { return ctx.c.GetHeader(key) }
func (ctx *GinContext) Body() []byte                                       { b, _ := io.ReadAll(ctx.c.Request.Body); return b }
func (ctx *GinContext) JSON(status int, data interface{})                  { ctx.c.JSON(status, data) }
func (ctx *GinContext) String(status int, data string)                     { ctx.c.String(status, data) }
func (ctx *GinContext) Status(status int)                                  { ctx.c.Status(status) }
func (ctx *GinContext) SetHeader(key, value string)                        { ctx.c.Header(key, value) }
func (ctx *GinContext) Next()                                              { ctx.c.Next() }
func (ctx *GinContext) Abort()                                             { ctx.c.Abort() }
func (ctx *GinContext) AbortWithStatus(status int)                         { ctx.c.AbortWithStatus(status) }
func (ctx *GinContext) AbortWithJSON(status int, data interface{})         { ctx.c.AbortWithStatusJSON(status, data) }
func (ctx *GinContext) Set(key string, value interface{})                  { ctx.c.Set(key, value) }
func (ctx *GinContext) Get(key string) (interface{}, bool)                 { return ctx.c.Get(key) }
