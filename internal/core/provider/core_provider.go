package provider

import (
	"sync"

	"cleanandclean/internal/core/interfaces"
)

var (
	instance *CoreProvider
	once     sync.Once
)

type ICoreProvider interface {
	interfaces.ICoreFactory
}

type CoreProvider struct {
	factory interfaces.ICoreFactory
}

func Init(factory interfaces.ICoreFactory) {
	once.Do(func() {
		instance = &CoreProvider{factory: factory}
	})
}

func Instance() ICoreProvider {
	if instance == nil {
		panic("CoreProvider not initialized. Call Init() first.")
	}
	return instance
}

func (p *CoreProvider) GetConfig() interfaces.IConfig {
	return p.factory.GetConfig()
}

func (p *CoreProvider) GetDatabase() interfaces.IDatabase {
	return p.factory.GetDatabase()
}

func (p *CoreProvider) GetServiceContainer() interfaces.IServiceContainer {
	return p.factory.GetServiceContainer()
}
