package cobra

import "github.com/Qianjiachen55/Nwfw55/framework/contract"

func (c *Command) MustMakeApp() contract.App {
	return c.GetContainer().MustMake(contract.AppKey).(contract.App)
}

func (c *Command) MustMakeKernel() contract.Kernel  {
	return c.GetContainer().MustMake(contract.KernelKey).(contract.Kernel)
}
