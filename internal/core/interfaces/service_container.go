package interfaces

type IServiceContainer interface {
	Set(name string, service interface{})
	Get(name string) interface{}
}
