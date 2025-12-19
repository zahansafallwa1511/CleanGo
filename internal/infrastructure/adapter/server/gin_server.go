package server

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"cleanandclean/internal/adapter/interfaces"
)

// GinServer implements IServer using Gin framework
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

func (s *GinServer) EnableRateLimiting(config interfaces.RateLimitConfig) {
	limiter := rate.NewLimiter(rate.Limit(config.RequestsPerMinute)/60, config.BurstSize)

	s.engine.Use(func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "RATE_LIMIT_EXCEEDED",
					"message": "Too many requests",
				},
			})
			return
		}
		c.Next()
	})
}

// GinRouteGroup implements IRouteGroup using Gin's RouterGroup
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
