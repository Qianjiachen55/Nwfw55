package framework

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	Bind(provide ServiceProvider) error

	IsBind(key string) bool

	Make(key string) (interface{}, error)

	MustMake(key string) interface{}

	MakeNew(key string,params []interface{}) (interface{},error)
}

type NwfwContainer struct {
	Container
	providers map[string]ServiceProvider

	instances map[string]interface{}

	lock sync.RWMutex
}

func NewNwfwContainer() *NwfwContainer {
	return &NwfwContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

func (nwfw *NwfwContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range nwfw.providers {
		name := provider.Name()
		line := fmt.Sprintf(name)

		ret = append(ret, line)
	}

	return ret
}

func (nwfw *NwfwContainer) Bind(provider ServiceProvider) error {
	nwfw.lock.Lock()
	//defer nwfw.lock.Unlock()

	key := provider.Name()
	nwfw.providers[key] = provider
	nwfw.lock.Unlock()
	if provider.IsDefer() == false {
		if err := provider.Boot(nwfw); err != nil {
			return err
		}

		params := provider.Params(nwfw)
		method := provider.Register(nwfw)
		instance, err := method(params...)

		if err != nil {
			return errors.New(err.Error())
		}

		nwfw.instances[key] = instance
	}

	return nil
}

func (nwfw *NwfwContainer) IsBind(key string) bool {
	return nwfw.findServiceProvider(key) != nil
}

func (nwfw *NwfwContainer) findServiceProvider(key string) ServiceProvider {
	nwfw.lock.RLock()
	defer nwfw.lock.RUnlock()

	if sp, ok := nwfw.providers[key]; ok {
		return sp
	}
	return nil
}

func (nwfw *NwfwContainer) Make(key string) (interface{}, error) {
	return nwfw.make(key, nil, false)
}

func (nwfw *NwfwContainer) MustMake(key string) interface{} {
	serv, err := nwfw.make(key, nil, false)
	if err != nil {
		panic(err)
	}

	return serv
}

func (nwfw *NwfwContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return nwfw.make(key, params, true)
}

func (nwfw *NwfwContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(nwfw); err != nil {
		return nil, err
	}

	if params == nil {
		params = sp.Params(nwfw)
	}
	method := sp.Register(nwfw)

	ins, err := method(params...)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return ins, err
}

// 实例化服务
func (nwfw *NwfwContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	nwfw.lock.RLock()
	defer nwfw.lock.RUnlock()

	sp := nwfw.findServiceProvider(key)

	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return nwfw.newInstance(sp, params)
	}

	if ins, ok := nwfw.instances[key]; ok {
		return ins, nil
	}

	inst, err := nwfw.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	nwfw.instances[key] = inst

	return inst, nil

}
