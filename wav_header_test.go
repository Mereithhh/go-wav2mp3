package wav2mp3_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mereithhh/go-wav2mp3"
)

func TestGetWavHeaderSize(t *testing.T) {

	inputBytes, _ := os.ReadFile("test2.wav")
	size := wav2mp3.GetWavHeaderSize(inputBytes)
	fmt.Println(size)
	t.Log(size)
	os.WriteFile("test_header2.txt", []byte(fmt.Sprintf("%d", size)), 0644)
}