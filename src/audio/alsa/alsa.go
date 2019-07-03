package alsa

import (
	"fmt"

	svhapi "github.com/DgINC/SVH4ALL/src/api"
	"github.com/yobert/alsa"
)

func getPlaybackDevices() ([]*alsa.Card, error) {
	cards, err := alsa.OpenCards()
	defer alsa.CloseCards(cards)
	if err != nil {
		svhapi.Logger.Info("logger construction succeeded")
		fmt.Println(err)
		return nil, err
	}
}
