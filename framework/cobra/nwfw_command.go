package cobra

import (
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/robfig/cron/v3"
	"log"
)

// SetContainer 设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// GetContainer 获取容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}

type CronSpec struct {
	Type string
	Cmd  *Command
	Spec string
	ServiceName string
}


func (c *Command) AddCronCommand(spec string, cmd *Command, args ...string) {
	root := c.Root()
	if root.Cron == nil {
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		//root.Cron = cron.New()
		root.CronSpecs = []CronSpec{}

	}
	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Cmd:  cmd,
		Spec: spec,
		Type: "normal-cron",
	})

	root.Cron.AddFunc(spec, func() {
		var cronCmd Command
		ctx := root.Context()
		cronCmd = *cmd
		cronCmd.SetParentNull()
		cronCmd.args = []string{}
		err := cronCmd.ExecuteContext(ctx)

		if err != nil{
			log.Println(err)
		}
	})

}

func (c *Command) SetParentNull() {
	c.parent = nil
}
