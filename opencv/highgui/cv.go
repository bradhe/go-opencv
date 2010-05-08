package highgui

/*
#include <opencv/cv.h>

//int depth_8U(){ return IPL_DEPTH_16U; }
*/
import "C"

//=========================================================================

type Point struct {
	X, Y	int
}

type Size struct {
	Width	int;
	Height	int
}

type Rect struct {
	X, Y	int
	Width	int
	Height	int
}

const IPL_DEPTH_8U = 1;//C.depth_8U();

const IPL_DEPTH_16U = 0;

// 图像结构
type Image struct {
	cptr	*C.IplImage
}

// 图像是否为空
func (img *Image)IsNull() bool {
	if img.cptr == nil { return true; }
	if img.cptr.imageData == nil { return true; }

	return false;
}

// 创建图形
func (img *Image)Create(size Size, depth, channels int) {
	if !img.IsNull() { img.Release() }

	img.cptr = C.cvCreateImage(size.toCvSize(), C.int(depth), C.int(channels));
}

//
func (s Size)toCvSize() C.CvSize {
	return C.cvSize(C.int(s.Width),C.int(s.Height));
}

func (img *Image)Release() {

}

func init_cv() {

	r := Rect{ X:10, Y:20, Width:20, Height: 30 }

	r.X++;

	//r.Y = C.depth_8U();

	C.cvPoint(C.int(0), C.int(0));
}
