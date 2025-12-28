package boot

import (
	"cleanandclean/internal/adapter/controllers"
	adapterInterfaces "cleanandclean/internal/adapter/interfaces"
	coreInterfaces "cleanandclean/internal/core/interfaces"
)

type Route struct {
	Method  string
	Path    string
	Handler interface{}
}

func Routes(c coreInterfaces.IServiceContainer) []Route {
	health := c.Get("HealthController").(*controllers.HealthController)
	post := c.Get("PostController").(*controllers.PostController)

	return []Route{
		{"GET", "/health", health.Check},

		{"POST", "/api/posts", post.Create},
		{"GET", "/api/posts", post.List},
		{"GET", "/api/posts/:id", post.Get},
		{"PUT", "/api/posts/:id", post.Update},
		{"DELETE", "/api/posts/:id", post.Delete},
	}
}

func RegisterRoutes(router adapterInterfaces.IRouter, routes []Route) {
	for _, r := range routes {
		switch r.Method {
		case "GET":
			router.GET(r.Path, r.Handler)
		case "POST":
			router.POST(r.Path, r.Handler)
		case "PUT":
			router.PUT(r.Path, r.Handler)
		case "DELETE":
			router.DELETE(r.Path, r.Handler)
		case "PATCH":
			router.PATCH(r.Path, r.Handler)
		}
	}
}
