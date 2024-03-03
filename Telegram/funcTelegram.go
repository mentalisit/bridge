package tg

import (
	"github.com/mentalisit/logger"
	"github.com/mentalisit/models"
	"strconv"
)

type Telegram struct {
	log *logger.Logger
}

func NewTelegram(log *logger.Logger) *Telegram {
	return &Telegram{log: log}
}

func (t *Telegram) DeleteMessage(ChatId string, MesId int) {
	var s models.ActionStruct
	s.Action = models.DeleteMessage
	s.Message = models.DeleteMessageStruct{
		MessageId: strconv.Itoa(MesId),
		Channel:   ChatId,
	}
	t.MarshalDataTelegram(s)
}
func (t *Telegram) SendChannelDelSecond(chatId, text string, second int) {
	var s models.ActionStruct
	s.Action = models.SendChannelDelSecond
	s.Message = models.SendTextDeleteSeconds{
		Text:    text,
		Channel: chatId,
		Seconds: second,
	}
	t.MarshalDataTelegram(s)
}
