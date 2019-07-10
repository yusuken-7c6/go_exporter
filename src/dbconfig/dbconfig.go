package dbconfig

import (
	"github.com/kelseyhightower/envconfig"
)

type DbEnv struct {
	Hostname string
	Password string
	Username string
	Name string
}


const env = &DbEnv{}
func Generate(prefix string) {
	envconfig.Process(prefix, &env)
}