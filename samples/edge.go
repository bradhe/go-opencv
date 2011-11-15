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

	w := opencv.GetSizeWidth(image)
	h := opencv.GetSizeHeight(image)

	// Create the output image
	cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3);

	// Convert to grayscale
	gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1);
	edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1);
	opencv.CvtColor(image, gray, opencv.CV_BGR2GRAY);

	var edge_thresh int = 1

	win := opencv.NewWindow("Edge",
		func(event, x, y, flags int, param interface{}) {
		//	fmt.Printf("event = %d, x = %d, y = %d, flags = %d\n", event, x, y, flags)	
		})

	win.CreateTrackbar("Edge", 1, 100, func(pos int) {
		//if pos > 0 { edge_thresh = pos }
		edge_thresh = pos

		opencv.Smooth( gray, edge, opencv.CV_BLUR, 3, 3, 0, 0 );
		opencv.Not( gray, edge );

		// Run the edge detector on grayscale
		opencv.Canny(gray, edge, float64(edge_thresh),  float64(edge_thresh*3), 3);

		opencv.Zero( cedge );
		// copy edge points
		opencv.Copy( image, cedge, edge );

		win.ShowImage(cedge);
		//cvShowImage(wndname, cedge);

		fmt.Printf("pos = %d\n", pos)
	})
	win.ShowImage(image);

	opencv.WaitKey(0)
}

