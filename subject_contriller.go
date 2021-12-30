package main

import (
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
	"github.com/Qianjiachen55/Nwfw55/provider/demo"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}


func SubjectListController(c *gin.Context) {
	demoService := c.MustMake(demo.Key).(demo.Service)

	foo := demoService.GetFoo()

	c.ISetOkStatus().IJson(foo)
}



func SubjectDelController(c *gin.Context)  {
	c.ISetOkStatus().IJson(" SubjectDelController")
}

func SubjectUpdateController(c *gin.Context)  {
	c.ISetOkStatus().IJson(" SubjectUpdateController")
}

func SubjectGetController(c *gin.Context)  {
	c.ISetOkStatus().IJson(" SubjectGetController")
}

func SubjectNameController(c *gin.Context)  {
	c.ISetOkStatus().IJson(" SubjectNameController")
}