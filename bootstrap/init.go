package bootstrap

import (
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/blacklist"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/env"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitApplication(blackListData []byte) {
	if env.SystemConfig.Debug {
		log.SetLevel(logrus.DebugLevel)
		gin.SetMode(gin.DebugMode)
	} else {
		log.SetLevel(logrus.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}

	log.Log().Info("Application initialized")

	blacklist.Init(blackListData)
}
