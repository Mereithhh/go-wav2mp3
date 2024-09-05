# go wav to mp3

简单封装了一下，将 wav 文件转换为 mp3 文件

## 依赖

先安装 `lame`:

### macos

```shell
brew install lame
```

**安完需要修改环境变量**

```shell
export CGO_CFLAGS="-I/opt/homebrew/opt/lame/include"
export CGO_LDFLAGS="-L/opt/homebrew/opt/lame/lib"
```

### linux

```shell
apt-get install libmp3lame-dev
```

### 库

```go
go get github.com/mereithhh/gowav2mp3
```

**用了 cgo，build 的时候记得打开**

## 使用方法

```go
package main

import (
	"fmt"
	"github.com/mereithhh/go-wav2mp3"
)

func main() {
  inputBytes, _ := os.ReadFile("test.wav")

	config := wav2mp3.NewWav2Mp3Config().SetBrate(128).SetChannel(1).SetInputSampleRate(22050).SetQuality(2)

	outputBytes := wav2mp3.Wav2Mp3(inputBytes, config)

	os.WriteFile("test.mp3", outputBytes, 0644)

}
```

## 参考

github.com/viert/go-lame
