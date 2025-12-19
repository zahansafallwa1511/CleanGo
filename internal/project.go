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

// ProjectConfig holds all configuration for the project
type ProjectConfig struct {
	CORS      *interfaces.CORSConfig
	RateLimit *interfaces.RateLimitConfig
}

type Project struct {
	Name        string
	App         *Application
	Logger      ILogger
	Factory     IDependencyFactory
	StartTime   time.Time
	Config      *ProjectConfig
	initialized bool
}

func NewProject(name string, factory IDependencyFactory, config *ProjectConfig) *Project {
	if config == nil {
		config = &ProjectConfig{}
	}
	project := &Project{
		Name:        name,
		Factory:     factory,
		Config:      config,
		initialized: false,
	}
	return project
}

func (project *Project) Initialize(appName, appVersion string) {
	project.App = NewApplication(appName, appVersion)
	project.Logger = project.Factory.CreateLogger()

	server := project.Factory.CreateServer()

	// Apply CORS config if provided
	if project.Config.CORS != nil {
		server.EnableCORS(*project.Config.CORS)
	}

	// Apply rate limiting config if provided
	if project.Config.RateLimit != nil {
		server.EnableRateLimiting(*project.Config.RateLimit)
	}

	// Register all routes
	RegisterRoutes(server)

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

