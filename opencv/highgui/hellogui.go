
package main

import (
	"opencv/highgui"
)

func main() {
	img, _ := highgui.LoadImage("../../example/images/lena.jpg")
	if img != nil {
		highgui.SaveImage("1.bmp", img)

		highgui.NamedWindow("win1", false)
		highgui.ShowImage("win1", img)
		highgui.WaitKey(0)
		highgui.DestroyWindow("win1")
	}
}
