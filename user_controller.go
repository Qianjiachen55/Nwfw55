package main

import (
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	foo,_ := c.DefaultQueryString("foo","def")

	time.Sleep(10*time.Second)

	c.ISetOkStatus().IJson("ok,foo: "+foo)

}