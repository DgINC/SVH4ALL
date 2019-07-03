package alsa

import (
	"bytes"
	"fmt"
	"log"

	"github.com/yobert/alsa"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func getPlaybackDevices() ([]*alsa.Card, error) {
	cards, err := alsa.OpenCards()
	defer alsa.CloseCards(cards)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
}
