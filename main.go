package main

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/app/console"
	"github.com/Qianjiachen55/Nwfw55/app/http"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/provider/app"
	"github.com/Qianjiachen55/Nwfw55/framework/provider/kernel"
)

func main()  {
	fmt.Println("begin!!")

	container := framework.NewNwfwContainer()
	n := app.NwfwAppProvider{}
	container.Bind(&n)

	if engine, err := http.NewHttpEngine(); err != nil{
		container.Bind(&kernel.NwfwKernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)

}
