package main

import (
	"github.com/Qianjiachen55/Nwfw55/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	foo,_ := c.QueryString("foo","def")

	time.Sleep(10*time.Second)

	c.SetOkStatus().Json("ok,foo: "+foo)

	return nil
}