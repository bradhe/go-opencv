package  highgui

import (
	"image"
	"os"
)


// 视频接口
type Video interface {
	Size()(width, height int)
	GrabFrame()(img image.Image, err os.Error)
}
