package server

import (
	"github.com/mentalisit/models"
	"strings"
)

func (b *Bridge) logic() {
	if b.in.Config == nil {
		b.in.Config = &models.BridgeConfig{}
	}

	if strings.HasPrefix(b.in.Text, ".") {
		b.Command()
		return
	} else {
		b.logicMessage()
	}

}
