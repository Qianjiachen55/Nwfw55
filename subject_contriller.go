package main

import (
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}


func SubjectListController(c *gin.Context) {
	c.ISetOkStatus().IJson(" SubjectListController")
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