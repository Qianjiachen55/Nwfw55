package command

import (
	"context"
	"github.com/Qianjiachen55/Nwfw55/framework/cobra"
	"github.com/Qianjiachen55/Nwfw55/framework/contract"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)



func initAppCommand() *cobra.Command {
	appCommand.AddCommand(appStartCommand)
	return appCommand
}


var appCommand = &cobra.Command{
	Use:                        "app",
	Short:                      "业务应用控制命令",
	Long:                       "业务应用控制：启动，关闭，重启，查询",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

var appStartCommand = &cobra.Command{
	Use:                        "start",
	Short:                      "启动web服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()

		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)

		core := kernelService.HttpEngine()

		server := &http.Server{
			Addr:              ":8888",
			Handler:           core,
		}


		go func() {
			server.ListenAndServe()
		}()
		quit := make(chan os.Signal)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM,syscall.SIGQUIT)

		<- quit

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil{
			log.Fatal("server shutdown: ",err)
		}

		return nil
	},
}