package env

import (
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/contract"
)

type NwfwEnvProvider struct {
	Folder string
}

func (provider *NwfwEnvProvider)  Register(c framework.Container) framework.NewInstance{
	return NewNwfwEnv
}


func (provider *NwfwEnvProvider) Boot (c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()

	return nil
}
func (provider *NwfwEnvProvider) IsDefer() bool {
	return false
}
func (provider *NwfwEnvProvider)  Params(c framework.Container) []interface{}{
	return []interface{}{provider.Folder}

}
func (provider *NwfwEnvProvider) Name() string {
	return contract.EnvKey
}