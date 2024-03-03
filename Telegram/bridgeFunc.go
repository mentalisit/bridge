package tg

import (
	"github.com/mentalisit/models"
	"strings"
	"sync"
)

func (t *Telegram) SendBridgeAsync(chatid []string, text string, fileURL string, resultChannel chan<- models.MessageTg, wg *sync.WaitGroup) {
	defer wg.Done()

	fileURL = strings.TrimSpace(fileURL)

	m := models.BridgeSendToMessenger{
		Text:      text,
		ChannelId: chatid,
		FileUrl:   fileURL,
	}
	ChatId, MessageId := t.MarshalDataTelegramSendFile(m, "file")
	resultChannel <- models.MessageTg{
		ChatId:    ChatId,
		MessageId: MessageId,
	}
}
