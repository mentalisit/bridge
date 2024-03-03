package ds

import (
	"github.com/mentalisit/models"
	"sync"
)

func (d *Discord) SendBridgeAsync(text, username string, channelID []string, fileURL, Avatar string, reply *models.BridgeMessageReply, resultChannel chan<- models.MessageDs, wg *sync.WaitGroup) {

	defer wg.Done()

	m := models.BridgeSendToMessenger{
		Text:      text,
		Sender:    username,
		ChannelId: channelID,
		Avatar:    Avatar,
		FileUrl:   fileURL,
		Reply:     reply,
	}

	mesarray := d.MarshalDataSendBridgeAsync(m)

	for _, ds := range mesarray {
		resultChannel <- ds
	}
}
