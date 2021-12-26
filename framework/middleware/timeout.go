package middleware

import (
	"context"
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
		finish := make(chan struct{},1)
		panicChan := make(chan interface{},1)

		durationCtx, cancel := context.WithTimeout(c.BaseContext(),d)

		defer cancel()

		go func() {
			defer func() {
				if p := recover();p != nil{
					panicChan <- p
				}
			}()
			c.Next()

			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			c.Json(500,"time out")
			log.Println(p)
		case <- finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			c.Json(500,"time out")

		}
		return nil
	}
}
