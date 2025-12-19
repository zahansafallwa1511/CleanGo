package interfaces

import "context"

// IServer defines the contract for HTTP server implementations
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

// IRouteGroup defines the contract for route grouping
type IRouteGroup interface {
	Use(middlewares ...HandlerFunc)
	RegisterRoute(method, path string, handlers ...HandlerFunc)
	Group(prefix string, middlewares ...HandlerFunc) IRouteGroup
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	RequestsPerMinute int
	BurstSize         int
}
