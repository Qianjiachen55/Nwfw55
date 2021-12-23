package main

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"net/http"
)

func main()  {
	fmt.Println("begin!!")
	core := framework.NewCore()

	registerRouter(core)

	server := &http.Server{
		Addr:              ":8888",
		Handler:          	core,
	}
	server.ListenAndServe()

}
