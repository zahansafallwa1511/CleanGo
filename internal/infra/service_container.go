package infra

import "sync"

type ServiceContainer struct {
	mu       sync.RWMutex
	services map[string]interface{}
}

func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{
		services: make(map[string]interface{}),
	}
}

func (c *ServiceContainer) Set(name string, service interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.services[name] = service
}

func (c *ServiceContainer) Get(name string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.services[name]
}
