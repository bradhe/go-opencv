// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"opencv"
	"os"
)

func main() {
	opencv.InitSystem(os.Args)

	image := opencv.LoadImage("./lena.jpg", opencv.CV_LOAD_IMAGE_COLOR)
	if image == nil {
		panic("LoadImage fail")
	}

	winName := "Go-OpenCV"
	opencv.NamedWindow(winName, opencv.CV_WINDOW_AUTOSIZE)
	opencv.ShowImage(winName, image);

	opencv.WaitKey(0)
}

