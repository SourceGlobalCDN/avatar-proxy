package main

import (
	_ "embed"

	"github.com/SourceGlobalCDN/avatar-proxy/bootstrap"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/env"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/log"
	"github.com/SourceGlobalCDN/avatar-proxy/router"
)

//go:embed blacklist.json
var blacklistData []byte

func init() {
	bootstrap.InitApplication(blacklistData)
}

func main() {
	r := router.InitRouter()

	log.Log().Infof("Starting application on %s", env.SystemConfig.Listen)
	err := r.Run(env.SystemConfig.Listen)
	if err != nil {
		log.Log().Panicf("Error starting application: %s", err)
		return
	}
}
