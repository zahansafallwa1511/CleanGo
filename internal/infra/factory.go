package infra

import (
	"sync"

	adapterInterfaces "cleanandclean/internal/adapter/interfaces"
	coreInterfaces "cleanandclean/internal/core/interfaces"
	"cleanandclean/internal/infra/adapters"
	"cleanandclean/internal/infra/config"
	"cleanandclean/internal/infra/database"
)

type Factory struct {
	mu        sync.RWMutex
	config    coreInterfaces.IConfig
	database  coreInterfaces.IDatabase
	router    adapterInterfaces.IRouter
	container *ServiceContainer
}

func NewFactory() (*Factory, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	return &Factory{
		config: cfg,
	}, nil
}

func MustNewFactory() *Factory {
	f, err := NewFactory()
	if err != nil {
		panic(err)
	}
	return f
}

func (f *Factory) CreateServiceContainer() *ServiceContainer {
	return NewServiceContainer()
}

func (f *Factory) GetConfig() coreInterfaces.IConfig {
	return f.config
}

func (f *Factory) GetDatabase() coreInterfaces.IDatabase {
	f.mu.RLock()
	if f.database != nil {
		defer f.mu.RUnlock()
		return f.database
	}
	f.mu.RUnlock()

	f.mu.Lock()
	defer f.mu.Unlock()
	if f.database == nil {
		f.database = database.NewPostgresDB(f.config.Database())
	}
	return f.database
}

func (f *Factory) GetServiceContainer() coreInterfaces.IServiceContainer {
	f.mu.RLock()
	if f.container != nil {
		defer f.mu.RUnlock()
		return f.container
	}
	f.mu.RUnlock()

	f.mu.Lock()
	defer f.mu.Unlock()
	if f.container == nil {
		f.container = f.CreateServiceContainer()
	}
	return f.container
}

func (f *Factory) CreateRouter() adapterInterfaces.IRouter {
	return adapters.NewGinRouter()
}

func (f *Factory) GetRouter() adapterInterfaces.IRouter {
	f.mu.RLock()
	if f.router != nil {
		defer f.mu.RUnlock()
		return f.router
	}
	f.mu.RUnlock()

	f.mu.Lock()
	defer f.mu.Unlock()
	if f.router == nil {
		f.router = f.CreateRouter()
	}
	return f.router
}
