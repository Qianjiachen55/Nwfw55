package gin

import (
	"context"
	"github.com/Qianjiachen55/Nwfw55/framework"
)

func (ctx *Context) BaseContext() context.Context  {
	return ctx.Request.Context()
}

func (engine *Engine) Bind(provider framework.ServiceProvider) error{
	return engine.container.Bind(provider)
}

func (engine *Engine)  IsBind(key string)bool{
	return engine.container.IsBind(key)
}

func (ctx *Context)Make(key string) (interface{},error) {
	return ctx.container.Make(key)
}

func (ctx *Context) MakeNew (key string,params []interface{})(interface{},error){
	return ctx.container.MakeNew(key,params)
}

func (ctx *Context)  MustMake(key string) interface{}{
	return ctx.container.MustMake(key)
}
