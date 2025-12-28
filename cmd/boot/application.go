package boot

import (
	"context"

	"cleanandclean/internal/adapter/interfaces"
)

type Application struct {
	router interfaces.IRouter
}

func NewApplication(router interfaces.IRouter) *Application {
	return &Application{router: router}
}

func (a *Application) Run(addr string) error {
	return a.router.Run(addr)
}

func (a *Application) Shutdown(ctx context.Context) error {
	return a.router.Shutdown(ctx)
}

func (a *Application) Close() error {
	return a.router.Close()
}
