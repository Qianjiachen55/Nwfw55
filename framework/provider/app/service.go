package app

import (
	"errors"
	"flag"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/util"
	"path/filepath"
)

type NwfwApp struct {
	container  framework.Container //服务容器
	baseFolder string              // 基础路径
}

func (n NwfwApp) Version() string {
	return "2022-01-01"
}

func (n NwfwApp) BaseFolder() string {
	if n.baseFolder != "" {
		return n.baseFolder
	}
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder 参数，默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	return util.GetExecDirectory()
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

	return &NwfwApp{
		container:  container,
		baseFolder: baseFolder,
	},nil
}