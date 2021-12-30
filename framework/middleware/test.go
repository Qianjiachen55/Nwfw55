package middleware

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
)

func Test1() gin.HandlerFunc  {


	return func(c *gin.Context)  {
		fmt.Println("pre test1")
		c.Next()
		fmt.Println("post test1")
	}
}

func Test2() gin.HandlerFunc {

	return func(c *gin.Context)  {
		fmt.Println("pre test2")
		c.Next()
		fmt.Println("post test2")
	}
}

func Test3() gin.HandlerFunc {
	return func(c *gin.Context)  {
		fmt.Println("pre test3")
		c.Next()
		fmt.Println("post test3")
	}
}