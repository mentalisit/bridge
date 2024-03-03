package tg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (t *Telegram) MarshalDataTelegram(message any) {
	data, err := json.Marshal(message)
	if err != nil {
		t.log.ErrorErr(err)
		return
	}

	_, err = http.Post("http://telegram/data", "application/json", bytes.NewBuffer(data))
	if err != nil {
		_, err = http.Post("http://192.168.100.155:803/data", "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Error sending message to telegram delete:", err)
			return
		}
	}
}

func (t *Telegram) MarshalDataTelegramSendFile(message any, key string) (string, string) {
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return "", ""
	}

	resp, err := http.Post(fmt.Sprintf("http://telegram/send/%s", key), "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error sending message to WebhookReplyDs:", err)
		return "", ""
	}
	var dataReply struct {
		MessageId string
		ChatId    string
	}
	err = json.NewDecoder(resp.Body).Decode(&dataReply)
	if err != nil {
		return "", ""
	}
	return dataReply.ChatId, dataReply.MessageId
}
