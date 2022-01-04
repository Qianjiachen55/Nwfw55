package command

import (
	"fmt"
	"github.com/Qianjiachen55/Nwfw55/framework/cobra"
	"github.com/Qianjiachen55/Nwfw55/framework/contract"
	"github.com/Qianjiachen55/Nwfw55/framework/util"
	"github.com/erikdubbelboer/gspt"
	"github.com/sevlyar/go-daemon"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var cronDaemon = false

func initCronCommand() *cobra.Command {

	cronStartCommand.Flags().BoolVarP(&cronDaemon, "daemon", "d", false, "start serve daemon")
	cronCommand.AddCommand(cronRestartCommand)
	cronCommand.AddCommand(cronListCommand)
	cronCommand.AddCommand(cronStateCommand)
	cronCommand.AddCommand(cronStopCommand)
	cronCommand.AddCommand(cronStartCommand)

	return cronCommand

}

var cronCommand = &cobra.Command{
	Use:   "cron",
	Short: "定时任务相关命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}
var cronRestartCommand = &cobra.Command{
	Use:   "restart",
	Short: "重启",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		servicePidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		content, err := ioutil.ReadFile(servicePidFile)

		if err != nil {
			return err
		}

		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))

			if err != nil {
				return err
			}
			if util.CheckProcessExist(pid) {
				if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
					return nil
				}
				for i := 0; i < 10; i++ {
					if util.CheckProcessExist(pid) == false {
						break
					}
					time.Sleep(1 * time.Second)
				}
				fmt.Println("kill process:" + strconv.Itoa(pid))
			}
		}
		cronDaemon = true
		return cronStartCommand.RunE(c, args)
	},
}

var cronListCommand = &cobra.Command{
	Use:   "list",
	Short: "list all schedule",
	RunE: func(c *cobra.Command, args []string) error {
		cronSpecs := c.Root().CronSpecs

		ps := [][]string{}
		for _, cronSpec := range cronSpecs {
			line := []string{cronSpec.Spec, cronSpec.Cmd.Use, cronSpec.Cmd.Short}
			ps = append(ps, line)
		}
		util.PrettyPrint(ps)

		return nil
	},
}

var cronStartCommand = &cobra.Command{
	Use:   "start",
	Short: "start cron daemon",
	RunE: func(c *cobra.Command, args []string) error {

		container := c.GetContainer()

		appService := container.MustMake(contract.AppKey).(contract.App)

		pidFolder := appService.RuntimeFolder()
		serverPidFile := filepath.Join(pidFolder, "cron.pid")
		logFolder := appService.LogFolder()
		serverLogFile := filepath.Join(logFolder, "cron.log")
		currentFold := util.GetExecDirectory()

		if cronDaemon {
			cntxt := &daemon.Context{
				PidFileName: serverPidFile,
				PidFilePerm: 0664,
				LogFileName: serverLogFile,
				LogFilePerm: 0640,
				WorkDir:     currentFold,
				Args:        []string{"", "cron", "start", "--daemon=true"},
				Umask:       027,
			}

			d, err := cntxt.Reborn()
			if err != nil {
				return err
			}
			if d != nil {
				fmt.Println("cron serve started, pid: ", d.Pid)
				fmt.Println("log file", serverLogFile)
				return nil
			}

			defer cntxt.Release()

			fmt.Println("daemon started")
			gspt.SetProcTitle("Nwfw cron")
			c.Root().Cron.Run()
			return nil
		}

		fmt.Println("start cron job")
		content := strconv.Itoa(os.Getpid())
		fmt.Println("[PID]", content)
		err := ioutil.WriteFile(serverPidFile, []byte(content), 0664)
		if err != nil {
			return err
		}

		gspt.SetProcTitle("nwfw cron")
		c.Root().Cron.Run()
		return nil
	},
}

var cronStopCommand = &cobra.Command{
	Use:   "stop",
	Short: "停止cron常驻进程",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		servicePidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		content, err := ioutil.ReadFile(servicePidFile)

		if err != nil {
			return err
		}

		if container != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
				return err
			}
			if err := ioutil.WriteFile(servicePidFile, []byte{}, 0644); err != nil {
				return err
			}
			fmt.Println("stop pid: ", pid)
		}
		return nil
	},
}

var cronStateCommand = &cobra.Command{
	Use:   "state",
	Short: "cron常驻进程状态",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		content, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			return err
		}
		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			if util.CheckProcessExist(pid) {
				fmt.Println("cron server started, pid:", pid)
				return nil
			}
		}
		fmt.Println("no cron server start")
		return nil
	},
}
