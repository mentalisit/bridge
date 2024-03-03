package main

import (
	"github.com/mentalisit/bridge/server"
	"github.com/mentalisit/logger"
)

func main() {
	cfg := InitConfig()

	log := logger.LoggerZapDiscord(cfg.Logger.Webhook)

	bridge := server.NewBridge(log)

	bridge.LoadConfig()

	log.Info("Bridge start")

	bridge.ServerRun()
}
