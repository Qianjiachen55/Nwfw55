package main

import (
	"context"
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"log"
	"time"
)

func FooControllerHandler(ctx *framework.Context)  error{
	fmt.Println(ctx.GetRequest().URL)
	finish := make(chan struct{},1)
	panicChan := make(chan interface{},1)


	//1. 生成超时控制context
	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(),time.Duration(1*time.Second))
	defer cancel()


	// 2. 	业务逻辑
	go func() {
		defer func() {
			if p:= recover();p != nil{
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)

		ctx.SetOkStatus().Json("ok")

		finish <- struct{}{}
	}()
	//3.设计事件处理顺序，当前 Goroutine 监听超时时间
				//Context 的 Done() 事件，
				//和具体的业务处理结束事件，
	//哪个先到就先处理哪个。
	select {
	case p := <-panicChan:
		// 业务逻辑出错
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()

		log.Println(p)
		ctx.SetStatus(500).Json("panic!")
	case <- finish:
		//正常处理结束
		fmt.Println("finish")
	case <- durationCtx.Done():
		//超时
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.SetStatus(500).Json("time out")
		ctx.SetHasTimeout()
	}
	return nil
}
