package wav2mp3

import (
	"bytes"
)

func GetWavHeaderSize(wavBytes []byte) int {
	// 读取一直到 data 区块，往后再加 4 个字节
	dataIndex := bytes.Index(wavBytes, []byte("data"))
	if dataIndex == -1 {
		return 44
	}
	return dataIndex + 8
}