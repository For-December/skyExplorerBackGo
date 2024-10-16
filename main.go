package main

import (
	"skyExplorerBack/src/constant/config"
	"skyExplorerBack/src/router"
	"skyExplorerBack/src/utils/logger"
)

func main() {

	logger.Info("start")

	if err := router.Routers().Run(":" + config.EnvCfg.ServerPort); err != nil {
		logger.Error("run server error: ", err)
		return
	}
}
