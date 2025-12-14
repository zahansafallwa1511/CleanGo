package internal

import (
	"context"
	"time"

	"cleanandclean/internal/adapter/interfaces"
)

type ILogger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type IDependencyFactory interface {
	CreateServer() interfaces.IServer
	CreateLogger() ILogger
}

type Project struct {
	Name        string
	App         *Application
	Logger      ILogger
	Factory     IDependencyFactory
	StartTime   time.Time
	Config      map[string]interface{}
	initialized bool
}

func NewProject(name string, factory IDependencyFactory) *Project {
	project := &Project{
		Name:        name,
		Factory:     factory,
		Config:      make(map[string]interface{}),
		initialized: false,
	}
	return project
}

func (project *Project) Initialize(appName, appVersion string) {
	project.App = NewApplication(appName, appVersion)
	project.Logger = project.Factory.CreateLogger()

	server := project.Factory.CreateServer()

	for _, route := range GetRoutes() {
		server.RegisterRoute(route.Method, route.Path, route.Handlers...)
	}

	project.App.SetServer(server)
	project.initialized = true
}

func (project *Project) Start() error {
	if !project.initialized {
		return nil
	}
	project.StartTime = time.Now()
	if project.Logger != nil {
		project.Logger.Info("Project %s starting", project.Name)
		project.Logger.Info("Server starting on %s", project.App.Server.Addr())
	}
	return project.App.Run()
}

func (project *Project) Stop(ctx context.Context) error {
	if project.Logger != nil {
		project.Logger.Info("Project %s stopping", project.Name)
	}
	return project.App.Shutdown(ctx)
}

func (project *Project) SetConfig(key string, value interface{}) {
	project.Config[key] = value
}

func (project *Project) GetConfig(key string) interface{} {
	return project.Config[key]
}
