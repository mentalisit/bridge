package ds

import (
	"github.com/mentalisit/logger"
	"github.com/mentalisit/models"
)

type Discord struct {
	log *logger.Logger
}

func NewDiscord(log *logger.Logger) *Discord {
	return &Discord{log: log}
}

func (d *Discord) DeleteMessageDs(ChatId, MesId string) {
	var s models.ActionStruct
	s.Action = models.DeleteMessage
	s.Message = models.DeleteMessageStruct{
		MessageId: MesId,
		Channel:   ChatId,
	}
	d.MarshalDataDiscord(s)
}
func (d *Discord) SendChannelDelSecondDs(chatId, text string, second int) {
	var s models.ActionStruct
	s.Action = models.SendChannelDelSecond
	s.Message = models.SendTextDeleteSeconds{
		Text:    text,
		Channel: chatId,
		Seconds: second,
	}
	d.MarshalDataDiscord(s)
}
