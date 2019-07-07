package alsa

import (
	"github.com/DgINC/SVH4ALL/src/api"
	"github.com/yobert/alsa"
)

var (
	record_device    *alsa.Device
	playback_device  *alsa.Device
	record_devices   []*alsa.Device
	playback_devices []*alsa.Device
)

func getPlaybackDevices() ([]*alsa.Device, error) {
	cards, err := alsa.OpenCards()
	defer alsa.CloseCards(cards)
	if err != nil {
		api.ApiLogger.Error("ALSA not found error: ", err)
		return nil, err
	}

	for _, card := range cards {
		devices, err := card.Devices()
		if err != nil {
			api.ApiLogger.Error("Devices not found error: ", err)
			return nil, err
		}
		for _, device := range devices {
			if device.Type != alsa.PCM {
				continue
			}
			if device.Play && playback_device != nil {
				playback_devices = append(playback_device, device)
			}
		}
	}
	return nil, playback_devices
}
