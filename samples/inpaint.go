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
	filename := "./fruits.jpg"
	if len(os.Args) == 2 { filename = os.Args[1] }
    
	img0 := opencv.LoadImage(filename, opencv.CV_LOAD_IMAGE_COLOR)
	if img0 == nil {
		panic("LoadImage fail")
	}
	defer opencv.ReleaseImage(img0)

	fmt.Printf("Hot keys: \n" +
		"\tESC - quit the program\n" +
		"\tr - restore the original image\n" +
		"\ti or ENTER - run inpainting algorithm\n" +
		"\t\t(before running it, paint something on the image)\n",
	)

	w := opencv.GetSizeWidth(img0)
	h := opencv.GetSizeHeight(img0)

	img := opencv.CloneImage( img0 )
	inpainted := opencv.CloneImage( img0 );
	inpaint_mask := opencv.CreateImage(w, h, 8, 1 )

	opencv.Zero( inpaint_mask )
	//opencv.Zero( inpainted )

	win := opencv.NewWindow("image", true)
	defer win.Destroy()

	prev_pt := opencv.Point{-1,-1};
	win.SetMouseCallback(func(event, x, y, flags int, win *opencv.Window) {
		if img == nil { os.Exit(0) }

		if event == opencv.CV_EVENT_LBUTTONUP ||
		  (flags & opencv.CV_EVENT_FLAG_LBUTTON) == 0 {
			prev_pt = opencv.Point{-1,-1}	
		} else if event == opencv.CV_EVENT_LBUTTONDOWN {
			prev_pt = opencv.Point{x,y}
		} else if event == opencv.CV_EVENT_MOUSEMOVE &&
			 (flags & opencv.CV_EVENT_FLAG_LBUTTON) != 0 {
			pt := opencv.Point{x,y}
			if prev_pt.X < 0 { prev_pt = pt }

			rgb := opencv.ScalarAll(255.0)
			opencv.Line( inpaint_mask, prev_pt, pt, rgb, 5, 8, 0 )
			opencv.Line( img, prev_pt, pt, rgb, 5, 8, 0 );
			prev_pt = pt

			win.ShowImage( img )
		}
	})
	win.ShowImage(img)
	opencv.WaitKey(0)

	win2 := opencv.NewWindow("inpainted image", true)
	defer win2.Destroy()

	for {
		key := opencv.WaitKey(20)
		if key == 27 {
			os.Exit(0)
		} else if key == int("r"[0]) {
			opencv.Zero( inpaint_mask )
			opencv.Copy( img0, img, nil )
			win.ShowImage( img )
		} else if key == int("i"[0]) || key == int("\n"[0]) {
			opencv.Inpaint( img, inpaint_mask, inpainted, 3,
				opencv.CV_INPAINT_TELEA,
			)
		}	
		win2.ShowImage(inpainted)
	}
	os.Exit(0)
}

