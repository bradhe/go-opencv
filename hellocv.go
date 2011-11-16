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

	image := opencv.LoadImage("./samples/lena.jpg", opencv.CV_LOAD_IMAGE_COLOR)
	if image == nil {
		panic("LoadImage fail")
	}

	win := opencv.NewWindow("Go-Opencv", true)
	win.ShowImage(image);

	opencv.WaitKey(0)
}

