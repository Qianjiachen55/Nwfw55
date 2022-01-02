package app

import (
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/contract"
)

type NwfwAppProvider struct {
	BaseFolder string
}

func (n *NwfwAppProvider) Register(container framework.Container) framework.NewInstance{
	return NewNwfwApp
}

func (n *NwfwAppProvider) Boot(container framework.Container) error {
	return nil
}

func (n *NwfwAppProvider) IsDefer() bool {
	return false
}


func (n *NwfwAppProvider)Params(container framework.Container) []interface{} {
	return []interface{}{container, n.BaseFolder}
}



func (n *NwfwAppProvider) Name() string {
	return contract.AppKey
}

