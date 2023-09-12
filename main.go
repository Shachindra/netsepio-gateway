package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/netsepio/netsepio-gateway/app"
	"github.com/netsepio/netsepio-gateway/config/envconfig"
	"github.com/netsepio/netsepio-gateway/util/pkg/logwrapper"
)

func main() {
	app.Init()
	logwrapper.Log.Info("Starting MCube NFT Gateway Platform")
	addr := fmt.Sprintf(":%d", envconfig.EnvVars.APP_PORT)
	err := app.GinApp.Run(addr)
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}
