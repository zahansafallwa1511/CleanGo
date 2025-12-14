package internal

import (
	"context"

	"cleanandclean/internal/adapter/interfaces"
)

type Application struct {
	Name    string
	Version string
	Server  interfaces.IServer
	running bool
}

func NewApplication(name, version string) *Application {
	return &Application{
		Name:    name,
		Version: version,
		running: false,
	}
}

func (application *Application) SetServer(server interfaces.IServer) {
	application.Server = server
}

func (application *Application) Run() error {
	if application.Server == nil {
		return nil
	}
	application.running = true
	return application.Server.Start()
}

func (application *Application) Shutdown(ctx context.Context) error {
	application.running = false
	if application.Server != nil {
		return application.Server.Stop(ctx)
	}
	return nil
}

func (application *Application) IsRunning() bool {
	return application.running
}
