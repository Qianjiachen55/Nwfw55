package demo

import (
	"github.com/Qianjiachen55/Nwfw55/framework/cobra"
	"log"
)

func InitFoo() *cobra.Command {
	FooCommand.AddCommand(Foo1Command)
	return FooCommand
}

var FooCommand = &cobra.Command{
	Use:     "foo",
	Short:   "foo 简要说明",
	Long:    "foo long",
	Aliases: []string{"fo", "f"},
	Example: "foo example",
	RunE: func(cmd *cobra.Command, args []string) error {
		//container := cmd.GetContainer()
		log.Println("execute foo command")
		return nil
	},
}

var Foo1Command = &cobra.Command{
	Use:                        "foo1",
	Aliases:                    []string{"fo1","f1"},
	Short:                      "foo1 short",
	Long:                       "foo1 long",
	Example:                    "foo1 example",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		log.Println(container)
		return nil
	},
}
