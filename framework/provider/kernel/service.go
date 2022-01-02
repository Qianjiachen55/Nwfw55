package kernel

import (
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
	"net/http"
)


// web 引擎
type NwfwKernelService struct {
	engine *gin.Engine
}

func NewNwfwKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)

	return &NwfwKernelService{engine: httpEngine}, nil
}

func (s NwfwKernelService)HttpEngine() http.Handler {
	return s.engine
}