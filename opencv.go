// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

import(
//	"image"
)

//-----------------------------------------------------------------------------
// CvPNG
//-----------------------------------------------------------------------------

// RNG.RandInt
func (rng *RNG)RandInt() uint {
	return RandInt(rng)
}
// RNG.RandReal
func (rng *RNG)RandReal() float64 {
	return RandReal(rng)
}

//-----------------------------------------------------------------------------
// IplImage
//-----------------------------------------------------------------------------

func NewIplImage(w, h, depth, channels int) *IplImage {
	return nil
}

//func (self *IplImage)ColorModel() image.ColorModel {
//	return image.RGBAColorModel
//}
//func (self *IplImage)Bounds() image.Rectangle {
//	return image.Rectangle{{0,0},{0,0}}
//}
//func (self *IplImage)At(x, y int) image.Color {
//	return image.RGBAColor{0,0,0,0}
//}

//func (self *IplImage)Set(x, y int, c Color) {
//	//return image.RGBAColor{0,0,0,0}
//}

func (self *IplImage)Width() int {
	return 0
}
func (self *IplImage)Height() int {
	return 0
}

//-----------------------------------------------------------------------------
// CvMat
//-----------------------------------------------------------------------------


//-----------------------------------------------------------------------------
// Window
//-----------------------------------------------------------------------------

// mouse callback
type MouseFunc func(event, x, y, flags int, param interface{})
// trackbar callback
type TrackbarFunc func(pos int)

// named window
type Window struct {
	name           string
	flags          int
	mouseHandle    MouseFunc
	param          interface{}
	trackbarHandle map[string]TrackbarFunc
	trackbarMax    map[string]int
	trackbarVal    map[string]int
	refCount       int
}

func NewWindow(name string, on_mouse MouseFunc) *Window {
	win := &Window{
		name:name, flags:CV_WINDOW_AUTOSIZE,
		mouseHandle:on_mouse, param:nil,
		trackbarHandle:make(map[string]TrackbarFunc, 50),
		trackbarMax:make(map[string]int, 50),
		trackbarVal:make(map[string]int, 50),
	}
	NamedWindow(name, CV_WINDOW_AUTOSIZE)
	SetMouseCallback(name)
	allWindows[name] = win
	return win
}
func (self *Window)CreateTrackbar(name string, value, count int, on_changed TrackbarFunc)  {
	self.trackbarVal[name] = value
	self.trackbarMax[name] = count
	self.trackbarHandle[name] = on_changed
	CreateTrackbar(name, self.name, value, count)
}

func (self *Window)Name() string {
	return self.name
}

func (self *Window)ShowImage(image *IplImage)  {
	ShowImage(self.name, image)
}
func (self *Window)Resize(width, height int)  {
	ResizeWindow(self.name, width, height)
}
func (self *Window)Move(x, y int)  {
	MoveWindow(self.name, x, y)
}
func (self *Window)Destroy()  {
	DestroyWindow(self.name)
	delete(allWindows, self.name)
}

//func (self *Window)GetTrackbarPos(name string) value, max int {
//	return self.trackbarVal[name], self.trackbarMax[name]
//}
func (self *Window)SetTrackbarPos(name string, value int) {
	SetTrackbarPos(name, self.name, value)
	self.trackbarVal[name] = value
}

//-----------------------------------------------------------------------------
// End
//-----------------------------------------------------------------------------

