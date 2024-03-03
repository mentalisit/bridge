package ds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mentalisit/models"
	"net/http"
)

func (d *Discord) MarshalDataDiscord(message any) {
	data, err := json.Marshal(message)
	if err != nil {
		d.log.ErrorErr(err)
		return
	}

	_, err = http.Post("http://discord/data", "application/json", bytes.NewBuffer(data))
	if err != nil {
		_, err = http.Post("http://192.168.100.155:802/data", "application/json", bytes.NewBuffer(data))
		d.log.ErrorErr(err)
		return
	}
}
func (d *Discord) MarshalDataSendBridgeAsync(message any) []models.MessageDs {
	data, err := json.Marshal(message)
	if err != nil {
		d.log.ErrorErr(err)
		return nil
	}

	resp, err := http.Post("http://discord/send/bridge", "application/json", bytes.NewBuffer(data))
	if err != nil {
		resp, err = http.Post("http://192.168.100.155:802/send/bridge", "application/json", bytes.NewBuffer(data))
		if err != nil {
			d.log.ErrorErr(err)
			return nil
		}
	}
	var dataReply []models.MessageDs
	err = json.NewDecoder(resp.Body).Decode(&dataReply)
	if err != nil {
		d.log.Info(fmt.Sprintf("err resp.Body %+v\n", resp.Body))
		d.log.ErrorErr(err)
		return nil
	}
	return dataReply
}
