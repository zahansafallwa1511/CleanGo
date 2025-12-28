package interfaces

type ICoreFactory interface {
	GetConfig() IConfig
	GetDatabase() IDatabase
	GetServiceContainer() IServiceContainer
}
