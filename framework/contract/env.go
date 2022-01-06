package contract

const (

	EnvProduction = "production"

	EnvTesting = "testing"

	EnvDevelopment = "development"

	EnvKey = "nwfw:env"
)

type Env interface {
	AppEnv() string

	IsExist( string) bool

	Get(string) string

	All() map[string]string

}

