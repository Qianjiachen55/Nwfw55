package command

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework/cobra"
	"github.com/Qianjiachen55/Nwfw55/framework/contract"
)

var DemoCommand = &cobra.Command{
	Use:                        "demo",
	Short:                      "demo for framework",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.KernelKey).(contract.App)
		fmt.Println("app base folder", appService.BaseFolder())
	},
}
