package dbconfig

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Env struct {
	Hostname string
	Password string
	Username string
	Name string
}

func Generate(prefix string) *Env {
	var env Env

	err := envconfig.Process(prefix, &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &env
}