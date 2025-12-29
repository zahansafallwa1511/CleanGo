package boot

import (
	"context"

	"cleanandclean/internal/adapter/controllers"
	coreInterfaces "cleanandclean/internal/core/interfaces"
	"cleanandclean/internal/core/provider"
	"cleanandclean/internal/core/usecases/post"
	"cleanandclean/internal/infra"
	"cleanandclean/internal/infra/repositories"
)

type Project struct {
	factory     *infra.Factory
	application *Application
}

func NewProject() *Project {
	factory := infra.MustNewFactory()

	// Initialize CoreProvider so use cases can access services
	provider.Init(infra.NewCoreFactoryAdapter(factory))

	container := factory.GetServiceContainer()
	router := factory.GetRouter()

	registerServices(container)

	routes := Routes(container)
	RegisterRoutes(router, routes)

	application := NewApplication(router)

	return &Project{
		factory:     factory,
		application: application,
	}
}

func registerServices(c coreInterfaces.IServiceContainer) {
	// Repositories
	c.Set("PostRepository", repositories.NewInMemoryPostRepository())

	// Use cases (no dependencies needed - they use CoreProvider)
	c.Set("CreatePostUseCase", post.NewCreatePostUseCase())
	c.Set("GetPostUseCase", post.NewGetPostUseCase())
	c.Set("ListPostsUseCase", post.NewListPostsUseCase())
	c.Set("UpdatePostUseCase", post.NewUpdatePostUseCase())
	c.Set("DeletePostUseCase", post.NewDeletePostUseCase())

	// Controllers
	c.Set("HealthController", controllers.NewHealthController())
	c.Set("PostController", controllers.NewPostController(
		c.Get("CreatePostUseCase").(*post.CreatePostUseCase),
		c.Get("GetPostUseCase").(*post.GetPostUseCase),
		c.Get("ListPostsUseCase").(*post.ListPostsUseCase),
		c.Get("UpdatePostUseCase").(*post.UpdatePostUseCase),
		c.Get("DeletePostUseCase").(*post.DeletePostUseCase),
	))
}

func (p *Project) Run(addr string) error {
	return p.application.Run(addr)
}

func (p *Project) Shutdown(ctx context.Context) error {
	return p.application.Shutdown(ctx)
}

func (p *Project) Close() error {
	return p.application.Close()
}

func (p *Project) Factory() *infra.Factory {
	return p.factory
}
