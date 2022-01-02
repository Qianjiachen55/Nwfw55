package kernel

import (
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/contract"
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
)

type NwfwKernelProvider struct {
	HttpEngine *gin.Engine
}

func (provider *NwfwKernelProvider) Register (c framework.Container) framework.NewInstance{
	return NewNwfwKernelService
}


func (provider *NwfwKernelProvider) Boot(c framework.Container) error{
	if provider.HttpEngine == nil{
		provider.HttpEngine = gin.Default()
	}

	provider.HttpEngine.SetContainer(c)
	return nil
}

func (provider *NwfwKernelProvider) IsDefer() bool {
	return false
}

func (provider *NwfwKernelProvider) Params (container framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

func (provider *NwfwKernelProvider) Name() string {
	return contract.KernelKey
}
