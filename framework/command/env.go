package command

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework/cobra"
	"github.com/Qianjiachen55/Nwfw55/framework/contract"
)

var envCommand = &cobra.Command{
	Use: "env",
	Short: "get current environment",
	Run: func(c *cobra.Command,args []string) {
		container := c.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)

		fmt.Println("enviroment: ", envService.AppEnv())
	},
}