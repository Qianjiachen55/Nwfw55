package app

import (
	"errors"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/util"
	"github.com/google/uuid"
	"path/filepath"
)

type NwfwApp struct {
	container  framework.Container //服务容器
	baseFolder string              // 基础路径
	appId	string //用于分布式服务

}




func (n NwfwApp) Version() string {
	return "2022-01-01"
}

func (n NwfwApp) BaseFolder() string {
	if n.baseFolder != "" {
		return n.baseFolder
	}

	return util.GetExecDirectory()
}


func (n NwfwApp) ConfigFolder() string {
	return filepath.Join(n.BaseFolder(),"config")
}

func (n NwfwApp) StorageFolder() string {
	return filepath.Join(n.BaseFolder(), "storage")
}

func (n NwfwApp) LogFolder() string {
	return filepath.Join(n.StorageFolder(), "log")
}

func (n NwfwApp) HttpFolder() string {
	return filepath.Join(n.BaseFolder(), "http")
}

func (n NwfwApp) ConsoleFolder() string {
	return filepath.Join(n.BaseFolder(), "console")
}

func (n NwfwApp) ProviderFolder() string {
	return filepath.Join(n.BaseFolder(), "provider")
}

func (n NwfwApp) MiddlewareFolder() string {
	return filepath.Join(n.BaseFolder(), "middleware")
}

func (n NwfwApp) CommandFolder() string {
	return filepath.Join(n.ConsoleFolder(), "command")
}

func (n NwfwApp) RuntimeFolder() string {
	return filepath.Join(n.BaseFolder(), "runtime")
}

func (n NwfwApp) TestFolder() string {
	return filepath.Join(n.BaseFolder(), "test")
}

func NewNwfwApp(params ...interface{}) (interface{},error) {
	if len(params) != 2{
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	//if baseFolder ==""{
	//	flag.StringVar(&baseFolder,"base_folder","","base_folder参数，默认为当前路径")
	//	flag.Parse()
	//}
	appId :=uuid.New().String()

	return &NwfwApp{
		appId: appId,
		container:  container,
		baseFolder: baseFolder,
	},nil
}

func (n NwfwApp) AppID() string {
	return n.appId
}