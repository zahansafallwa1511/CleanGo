package infra

import (
	coreinterfaces "cleanandclean/internal/core/interfaces"
)

var _ coreinterfaces.ICoreFactory = (*CoreFactoryAdapter)(nil)

type CoreFactoryAdapter struct {
	factory *Factory
}

func NewCoreFactoryAdapter(factory *Factory) *CoreFactoryAdapter {
	return &CoreFactoryAdapter{factory: factory}
}

func (a *CoreFactoryAdapter) GetConfig() coreinterfaces.IConfig {
	return a.factory.GetConfig()
}

func (a *CoreFactoryAdapter) GetDatabase() coreinterfaces.IDatabase {
	return a.factory.GetDatabase()
}

func (a *CoreFactoryAdapter) GetServiceContainer() coreinterfaces.IServiceContainer {
	return a.factory.GetServiceContainer()
}
