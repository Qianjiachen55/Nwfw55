package main

import (
	"context"
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
	"github.com/Qianjiachen55/Nwfw55/provider/demo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	fmt.Println("begin!!")

	core := gin.New()

	core.Bind(&demo.DemoServiceProvider{})

	core.Use(
		gin.Recovery(),
		)


	registerRouter(core)

	server := &http.Server{
		Addr:              ":8888",
		Handler:          	core,
	}

	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM,syscall.SIGQUIT)

	<- quit

	timeoutCtx,cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(timeoutCtx);err != nil{
		log.Fatal("Server Shutdown:", err)
	}


}
