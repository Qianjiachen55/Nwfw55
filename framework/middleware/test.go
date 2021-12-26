package middleware

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework"
)

func Test1() framework.ControllerHandler  {


	return func(c *framework.Context) error {
		fmt.Println("pre test1")
		c.Next()
		fmt.Println("post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {

	return func(c *framework.Context) error {
		fmt.Println("pre test2")
		c.Next()
		fmt.Println("post test2")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("pre test3")
		c.Next()
		fmt.Println("post test3")
		return nil
	}
}