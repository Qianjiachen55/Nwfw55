package contract

import "net/http"

const KernelKey = "Nwfw:kernel"

type Kernel interface {
	HttpEngine() http.Handler
}