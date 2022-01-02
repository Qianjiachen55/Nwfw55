package console

import (
	"github.com/Qianjiachen55/Nwfw55/app/console/command/demo"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/command"
	"github.com/Qianjiachen55/Nwfw55/framework/cobra"
)

func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:                        "nwfw",
		Short:                      "nwfw 命令",
		Long:                       "nwfw 提供命令行工具",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions:          cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	rootCmd.SetContainer(container)

	command.AddKernelCommands(rootCmd)

	AddAppCommand(rootCmd)

	return rootCmd.Execute()
}

func AddAppCommand(rootCmd *cobra.Command)  {

	rootCmd.AddCommand(demo.InitFoo())
}