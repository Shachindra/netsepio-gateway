package envconfig

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type config struct {
	APP_PORT       int      `env:"APP_PORT,required"`
	APP_NAME       string   `env:"APP_NAME,required"`
	GIN_MODE       string   `env:"GIN_MODE,required"`
	ALLOWED_ORIGIN []string `env:"ALLOWED_ORIGIN,required" envSeparator:","`
}

var EnvVars config = config{}

func InitEnvVars() {

	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
}
