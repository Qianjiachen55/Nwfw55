package gin

import "github.com/Qianjiachen55/Nwfw55/framework/contract"

func (c *Context)  MustMakeApp() contract.App{
	return c.MustMake(contract.AppKey).(contract.App)
}

func (c *Context) MustMakeKernel () contract.Kernel{
	return c.MustMake(contract.KernelKey).(contract.Kernel)
}
