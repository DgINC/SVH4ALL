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

type Sink struct {
	alsaPlaybackDevice *alsa.Device
}

func getPlaybackDevices() (playback_devices []*alsa.Device, err error) {
	cards, err := alsa.OpenCards()
	defer alsa.CloseCards(cards)
	if err != nil {
		api.ApiLogger.Error("ALSA not found")
		return nil, err
	}

	for _, card := range cards {
		devices, err := card.Devices()
		if err != nil {
			api.ApiLogger.Error("Devices not found")
			return nil, err
		}
		for _, device := range devices {
			if device.Type != alsa.PCM {
				continue
			}
			if device.Play && playback_device != nil {
				playback_devices = append(playback_devices, device)
			}
		}
	}
	return playback_devices, nil
}

func NewSink() *Sink {
	return &Sink{}
}

func (s *Sink) Sink(sourceID string, sampleRate, numChannels, bufferSize int) (func([][]float64) error, error) {

}

func (s *Sink) Flush(string) error {
	if s.stream == nil {
		return nil
	}
	err := s.stream.Stop()
	if err != nil {
		return err
	}
	err = s.stream.Close()
	if err != nil {
		return err
	}
	return portaudio.Terminate()
}
