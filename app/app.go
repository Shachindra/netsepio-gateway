package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/netsepio/netsepio-gateway/api"
	"github.com/netsepio/netsepio-gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
	"github.com/netsepio/netsepio-gateway/config/envconfig"
)

var GinApp *gin.Engine

func Init() {
	envconfig.InitEnvVars()
	gin.SetMode(envconfig.EnvVars.GIN_MODE)
	logwrapper.Init()
	GinApp = gin.Default()

	corsConfig := cors.New(cors.Config{AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     envconfig.EnvVars.ALLOWED_ORIGIN})
	GinApp.Use(corsConfig)
	api.ApplyRoutes(GinApp)
}
