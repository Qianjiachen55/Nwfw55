package contract

const AppKey = "Nwfw:app"

type App interface {

	Version() string

	BaseFolder() string

	ConfigFolder() string

	LogFolder() string

	ProviderFolder() string

	MiddlewareFolder() string

	CommandFolder() string

	RuntimeFolder() string

	TestFolder() string
}
