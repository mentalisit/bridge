package main

import (
	"github.com/mentalisit/logger"
	"github.com/mentalisit/rsbot/bridge/server"
)

func main() {
	cfg := InitConfig()

	log := logger.LoggerZapDiscord(cfg.Logger.Webhook)

	bridge := server.NewBridge(log)

	bridge.LoadConfig()

	log.Info("Bridge start")

	bridge.ServerRun()
}
