// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"fmt"
	"opencv"
)

func main() {
	filename := "./samples/lena.jpg"
	if len(os.Args) == 2 { filename = os.Args[1] }

	image := opencv.LoadImage(filename, opencv.CV_LOAD_IMAGE_COLOR)
	if image == nil {
		panic("LoadImage fail")
	}
	defer image.Release()

	win := opencv.NewWindow("Go-Opencv", true)
	defer win.Destroy()

	win.SetMouseCallback(func(event, x, y, flags int, win *opencv.Window) {
		fmt.Printf("event = %d, x = %d, y = %d, flags = %d\n", event, x, y, flags)      
	})
	win.CreateTrackbar("Thresh", 1, 100, func(pos int, win *opencv.Window) {
		fmt.Printf("pos = %d\n", pos)
	})
	
	win.ShowImage(image);

	opencv.WaitKey(0)
}

