package interfaces

import "context"

type HandlerFunc func(ctx IContext)

type IServer interface {
	Start() error
	Stop(ctx context.Context) error
	Addr() string
	Use(middlewares ...HandlerFunc)
	RegisterRoute(method, path string, handlers ...HandlerFunc)
	Group(prefix string, middlewares ...HandlerFunc) IRouteGroup
	EnableCORS(config CORSConfig)
	EnableRateLimiting(config RateLimitConfig)
}

type IRouteGroup interface {
	Use(middlewares ...HandlerFunc)
	RegisterRoute(method, path string, handlers ...HandlerFunc)
	Group(prefix string, middlewares ...HandlerFunc) IRouteGroup
}

type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
}

type RateLimitConfig struct {
	RequestsPerMinute int
	BurstSize         int
}
