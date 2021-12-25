package main

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"time"
)

func registerRouter(core *framework.Core)  {
	fmt.Println("register")
	core.Get("/user/login",framework.TimeoutHandler(UserLoginController,time.Second))


	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id",SubjectDelController)
		subjectApi.Put("/:id",SubjectUpdateController)
		subjectApi.Get("/:id",SubjectGetController)
		subjectApi.Get("/list/all",SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name",SubjectNameController)
		}
	}
}
