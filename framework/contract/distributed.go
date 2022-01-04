package contract

import "time"

const DistributeKey = "nwfw:distribute"

type Distributed interface {

	Select(serviceName string,appId string,holdTime time.Duration)(selectAppId string,err error)
}
