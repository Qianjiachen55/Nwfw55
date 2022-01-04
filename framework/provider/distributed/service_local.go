package distributed

import (
	"errors"
	"github.com/Qianjiachen55/Nwfw55/framework"
	"github.com/Qianjiachen55/Nwfw55/framework/contract"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

type LocalDistributedService struct {
	container framework.Container
}

func NewLocalDistributedService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	return &LocalDistributedService{container: container}, nil
}

func (s LocalDistributedService) Select(serviceName string, appID string, holdTime time.Duration) (selectAppID string, err error) {
	appService := s.container.MustMake(contract.AppKey).(contract.App)
	runtimeFolder := appService.RuntimeFolder()
	lockFile := filepath.Join(runtimeFolder, "distribute_"+serviceName)

	lock, err := os.OpenFile(lockFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	err = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil { // 没有抢到文件锁

		selectAppIDByt, err := ioutil.ReadAll(lock)
		if err != nil {
			return "", err
		}
		return string(selectAppIDByt), err
	}
	go func() {
		defer func() {

			syscall.Flock(int(lock.Fd()), syscall.LOCK_UN) // 释放文件锁

			lock.Close()

			os.Remove(lockFile)
		}()

		timer := time.NewTimer(holdTime) //计时器

		<-timer.C
	}()

	//抢到锁
	if _, err := lock.WriteString(appID); err != nil {
		return "", err
	}
	return appID, nil
}
