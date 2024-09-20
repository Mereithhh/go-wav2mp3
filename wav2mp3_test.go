package wav2mp3_test

import (
	"os"
	"testing"

	"github.com/mereithhh/go-wav2mp3"
)

func TestWav2Mp3(t *testing.T) {
	inputBytes, _ := os.ReadFile("test2.wav")

	config := wav2mp3.NewWav2Mp3Config().SetBrate(128).SetChannel(1).SetInputSampleRate(22050).SetQuality(2)

	outputBytes := wav2mp3.Wav2Mp3(inputBytes, config)

	os.WriteFile("test2.mp3", outputBytes, 0644)

}
