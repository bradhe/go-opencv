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


//-----------------------------------------------------------------------------
// End
//-----------------------------------------------------------------------------

