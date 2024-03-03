package server

import (
	"fmt"
	"github.com/mentalisit/models"
)

func (b *Bridge) LoadConfig() {
	var i = 0
	var bridge string
	bc, _ := GetBridgeConfig()
	for _, configBridge := range bc {
		b.configs[configBridge.NameRelay] = configBridge
		i++
		bridge = bridge + fmt.Sprintf("%s, ", configBridge.NameRelay)
	}
	fmt.Printf("Загружено конфиг мостов %d : %s\n", i, bridge)
}
func (b *Bridge) CacheNameBridge(nameRelay string) (bool, models.BridgeConfig) {
	if len(b.configs) != 0 {
		for _, config := range b.configs {
			if config.NameRelay == nameRelay {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
func (b *Bridge) AddNewBridgeConfig(br models.BridgeConfig) {
	b.configs[br.NameRelay] = br
	b.InsertBridgeChat(br)
}
func (b *Bridge) AddBridgeConfig(br models.BridgeConfig) {
	a := b.configs[br.NameRelay]
	if len(br.ChannelDs) > 0 {
		a.ChannelDs = append(a.ChannelDs, br.ChannelDs...)
	}
	if len(br.ChannelTg) > 0 {
		a.ChannelTg = append(a.ChannelTg, br.ChannelTg...)
	}
	b.UpdateBridgeChat(a)
	b.configs[br.NameRelay] = a
}
func (b *Bridge) InsertBridgeChat(br models.BridgeConfig) {
	MarshalDataStorage(br, "insert")
}
func (b *Bridge) UpdateBridgeChat(br models.BridgeConfig) {
	MarshalDataStorage(br, "update")
}

func (b *Bridge) CacheCheckChannelConfigDS(chatIdDs string) (bool, models.BridgeConfig) {
	for _, config := range b.configs {
		for _, ds := range config.ChannelDs {
			if ds.ChannelId == chatIdDs {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
func (b *Bridge) CacheCheckChannelConfigTg(chatIdTg string) (bool, models.BridgeConfig) {
	for _, config := range b.configs {
		for _, tg := range config.ChannelTg {
			if tg.ChannelId == chatIdTg {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
