package main

import "github.com/Qianjiachen55/Nwfw55/framework"

func SubjectAddController(c *framework.Context) error{
	c.SetOkStatus().Json("ok, SubjectAddController")
	return nil
}


func SubjectListController(c *framework.Context) error{
	c.SetOkStatus().Json(" SubjectListController")
	return nil
}



func SubjectDelController(c *framework.Context) error {
	c.SetOkStatus().Json(" SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.SetOkStatus().Json(" SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.SetOkStatus().Json(" SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.SetOkStatus().Json(" SubjectNameController")
	return nil
}