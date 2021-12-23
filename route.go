package main

import (
	"github.com/Qianjiachen55/Nwfw55/framework"
)

func registerRouter(core *framework.Core)  {
	//fmt.Println("register")
	core.Get("foo",FooControllerHandler)
}
