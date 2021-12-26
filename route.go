package main

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/middleware"
)

func registerRouter(core *framework.Core)  {
	fmt.Println("register")
	core.Get("/user/login",middleware.Test3(),UserLoginController)


	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id",SubjectDelController)
		subjectApi.Put("/:id",SubjectUpdateController)
		subjectApi.Get("/:id",middleware.Test3(),SubjectGetController)
		subjectApi.Get("/list/all",SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name",SubjectNameController)
		}
	}
}
