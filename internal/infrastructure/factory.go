package infrastructure

import (
	"cleanandclean/internal"
	"cleanandclean/internal/adapter/interfaces"
	"cleanandclean/internal/infrastructure/adapter/logger"
	"cleanandclean/internal/infrastructure/adapter/server"
)

type DefaultFactory struct{}

func NewDefaultFactory() *DefaultFactory {
	return &DefaultFactory{}
}

func (f *DefaultFactory) CreateServer() interfaces.IServer {
	return server.NewGinServer(":8080")
}

func (f *DefaultFactory) CreateLogger() internal.ILogger {
	return logger.NewDefaultLogger()
}
