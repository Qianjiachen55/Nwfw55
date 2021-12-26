package main

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/middleware"
	"net/http"
)

func main()  {
	fmt.Println("begin!!")
	core := framework.NewCore()
	core.Use(
		middleware.Test1(),
		middleware.Test2())


	registerRouter(core)

	server := &http.Server{
		Addr:              ":8888",
		Handler:          	core,
	}
	server.ListenAndServe()

}
