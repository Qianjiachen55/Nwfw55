package http

import "github.com/Qianjiachen55/Nwfw55/framework/gin"

func NewHttpEngine() (*gin.Engine,error) {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	Routes(r)

	return r,nil

}
