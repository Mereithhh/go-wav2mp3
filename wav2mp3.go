package wav2mp3

import (
	"bufio"
	"bytes"

	"github.com/viert/go-lame"
)

type Wav2Mp3Config struct {
	InputSampleRate int
	Channel         int
	Quality         int
	Brate           int
}

func NewWav2Mp3Config() *Wav2Mp3Config {
	return &Wav2Mp3Config{
		InputSampleRate: 22050,
		Channel:         1,
		Quality:         2,
		Brate:           128,
	}
}

func (c *Wav2Mp3Config) SetInputSampleRate(rate int) *Wav2Mp3Config {
	c.InputSampleRate = rate
	return c
}

func (c *Wav2Mp3Config) SetChannel(channel int) *Wav2Mp3Config {
	c.Channel = channel
	return c
}

func (c *Wav2Mp3Config) SetQuality(quality int) *Wav2Mp3Config {
	c.Quality = quality
	return c
}

func (c *Wav2Mp3Config) SetBrate(brate int) *Wav2Mp3Config {
	c.Brate = brate
	return c
}

func Wav2Mp3(wavBytes []byte, config *Wav2Mp3Config) (mp3bytes []byte) {
	headerSize := GetWavHeaderSize(wavBytes)
	wavBytes = wavBytes[headerSize:] // remove header
	byteWriter := &bytes.Buffer{}
	enc := lame.NewEncoder(byteWriter)
	enc.SetNumChannels(config.Channel)
	enc.SetQuality(config.Quality)
	enc.SetBrate(config.Brate)
	enc.SetInSamplerate(config.InputSampleRate)
	defer enc.Close()

	inputReader := bytes.NewReader(wavBytes)

	r := bufio.NewReader(inputReader)
	r.WriteTo(enc)

	return byteWriter.Bytes()
}
