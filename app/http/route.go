package http

import (
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
	"github.com/Qianjiachen55/Nwfw55/app/http/module/demo"
)

func Routes(r *gin.Engine)  {
	r.Static("/dist/","./dist")

	demo.Register(r)
}
