package alsa

import (
	"fmt"

	svhapi "github.com/DgINC/SVH4ALL/src/api"
	"github.com/yobert/alsa"
)

var (
	record_device   *alsa.Device
	playback_device *alsa.Device
)

func getPlaybackDevices() ([]*alsa.Card, error) {
	cards, err := alsa.OpenCards()
	defer alsa.CloseCards(cards)
	if err != nil {
		svhapi.APILogger.Info("logger construction succeeded")
		fmt.Println(err)
		return nil, err
	}
}
