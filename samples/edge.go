// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"opencv"
	"os"
	"fmt"
)

func main() {
	opencv.InitSystem(os.Args)

	image := opencv.LoadImage("./fruits.jpg", opencv.CV_LOAD_IMAGE_COLOR)
	if image == nil {
		panic("LoadImage fail")
	}
	defer opencv.ReleaseImage(image)

	w := opencv.GetSizeWidth(image)
	h := opencv.GetSizeHeight(image)

	// Create the output image
	cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3);
	defer opencv.ReleaseImage(cedge)

	// Convert to grayscale
	gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1);
	edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1);
	defer opencv.ReleaseImage(gray)
	defer opencv.ReleaseImage(edge)

	opencv.CvtColor(image, gray, opencv.CV_BGR2GRAY);

	win := opencv.NewWindow("Edge", true)
	defer win.Destroy()

	win.SetMouseCallback(func(event, x, y, flags int, win *opencv.Window) {
		fmt.Printf("event = %d, x = %d, y = %d, flags = %d\n", event, x, y, flags)	
	})

	win.CreateTrackbar("Thresh", 1, 100, func(pos int, win *opencv.Window) {
		edge_thresh := pos

		opencv.Smooth( gray, edge, opencv.CV_BLUR, 3, 3, 0, 0 );
		opencv.Not( gray, edge );

		// Run the edge detector on grayscale
		opencv.Canny(gray, edge, float64(edge_thresh),  float64(edge_thresh*3), 3);

		opencv.Zero( cedge );
		// copy edge points
		opencv.Copy( image, cedge, edge );

		win.ShowImage(cedge);

		fmt.Printf("pos = %d\n", pos)
	})
	win.ShowImage(image);

	for {
		key := opencv.WaitKey(20)
		if key == int("q"[0]) { os.Exit(0) }
	}

	os.Exit(0)
}

