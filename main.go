package main

import (
	"github.com/Qianjiachen55/Nwfw55/app/console"
	"github.com/Qianjiachen55/Nwfw55/app/http"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/provider/app"
	"github.com/Qianjiachen55/Nwfw55/framework/provider/distributed"
	"github.com/Qianjiachen55/Nwfw55/framework/provider/env"
	"github.com/Qianjiachen55/Nwfw55/framework/provider/kernel"
)

func main()  {
	//fmt.Println("begin!!")

	container := framework.NewNwfwContainer()

	container.Bind(&app.NwfwAppProvider{})



	container.Bind(&env.NwfwEnvProvider{})

	container.Bind(&distributed.LocalDistributedProvider{})
	//container.Bind(&distributed.LocalDistributedProvider{})

	if engine, err := http.NewHttpEngine(); err == nil{
		container.Bind(&kernel.NwfwKernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)

}
