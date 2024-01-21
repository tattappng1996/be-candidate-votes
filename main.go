package main

import (
	"be-cadidate-votes/appconfig"
	"be-cadidate-votes/routes"
	"be-cadidate-votes/utility/config"
	"be-cadidate-votes/utility/logger"
	"be-cadidate-votes/utility/timezone"
)

func main() {
	// setup logger
	l := logger.Init()

	// setup timezone
	timezone.SetEnvTimezone("Asia/Bangkok")
	cfg := config.LoadFileConfig[appconfig.AppConfig]()

	// setup route and server
	e := routes.ServerStart(cfg, l)
	routes.ServerShutdown(e)
}
