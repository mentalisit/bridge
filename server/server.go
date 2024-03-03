package server

import (
	"github.com/gin-gonic/gin"
	ds "github.com/mentalisit/bridge/Discord"
	tg "github.com/mentalisit/bridge/Telegram"
	"github.com/mentalisit/logger"
	"github.com/mentalisit/models"
	"net/http"
	"os"
)

type Bridge struct {
	log      *logger.Logger
	in       models.ToBridgeMessage
	messages []models.BridgeTempMemory
	configs  map[string]models.BridgeConfig
	discord  *ds.Discord
	telegram *tg.Telegram
}

func NewBridge(log *logger.Logger) *Bridge {
	bridge := &Bridge{
		log:      log,
		configs:  make(map[string]models.BridgeConfig),
		discord:  ds.NewDiscord(log),
		telegram: tg.NewTelegram(log),
	}
	return bridge
}

func (b *Bridge) ServerRun() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// Обработчик для принятия сообщений от DiscordService
	router.POST("/inbox/bridge", b.indoxBridge)

	err := router.Run(":80")
	if err != nil {
		b.log.ErrorErr(err)
		os.Exit(1)
	}
}
func (b *Bridge) indoxBridge(c *gin.Context) {
	var mes models.ToBridgeMessage

	if err := c.ShouldBindJSON(&mes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b.in = mes
	b.logic()

	c.JSON(http.StatusOK, gin.H{"message": "Message received successfully"})
}
