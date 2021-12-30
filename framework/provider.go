package framework

type NewInstance func(...interface{}) (interface{}, error)

type ServiceProvider interface {
	Register(Container) NewInstance

	// 实例化服务时调用，设计基础配置
	Boot(Container) error

	// 是否延迟实例化
	IsDefer() bool

	Params(Container) []interface{}

	Name() string
}
