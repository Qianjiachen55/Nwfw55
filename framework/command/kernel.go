package command

import "github.com/Qianjiachen55/Nwfw55/framework/cobra"

func AddKernelCommands(root *cobra.Command)  {
	root.AddCommand(DemoCommand)

	root.AddCommand(initAppCommand())

}
