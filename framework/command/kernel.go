package command

import "github.com/Qianjiachen55/Nwfw55/framework/cobra"

func AddKernelCommands(root *cobra.Command)  {
	root.AddCommand(initCronCommand())

	root.AddCommand(initAppCommand())
	root.AddCommand(envCommand)

}
