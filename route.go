package main

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
	"github.com/Qianjiachen55/Nwfw55/framework/middleware"
)

func registerRouter(core *gin.Engine)  {
	fmt.Println("register")
	core.GET("/user/login",middleware.Test3(),UserLoginController)


	subjectApi := core.Group("/subject")
	{
		subjectApi.DELETE("/:id",SubjectDelController)
		subjectApi.PUT("/:id",SubjectUpdateController)
		subjectApi.GET("/:id",middleware.Test3(),SubjectGetController)
		subjectApi.DELETE("/list/all",SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name",SubjectNameController)
		}
	}
}
