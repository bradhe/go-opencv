// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

/*
#cgo LDFLAGS: -lcxcore

#include <opencv/cxcore.h>
#include <stdlib.h>
#include <string.h>

*/
import "C"
import (
	//"errors"
	"unsafe"
)

func init() {
}

/****************************************************************************************\
* Array allocation, deallocation, initialization and access to elements       *
\****************************************************************************************/

func Alloc(size int) unsafe.Pointer {
	return unsafe.Pointer(C.cvAlloc(C.size_t(size)))
}
func Free(p unsafe.Pointer) {
	C.cvFree_(p)
}


/* Allocates and initializes IplImage header */
//CVAPI(IplImage*)  cvCreateImageHeader( CvSize size, int depth, int channels );

/* Inializes IplImage header */
//CVAPI(IplImage*) cvInitImageHeader( IplImage* image, CvSize size, int depth,
//                                   int channels, int origin CV_DEFAULT(0),
//                                   int align CV_DEFAULT(4));

/* Creates IPL image (header and data) */
func CreateImage(w, h, depth, channels int) *IplImage {
	size := C.cvSize(C.int(w), C.int(h))
	img := C.cvCreateImage(size, C.int(depth), C.int(channels))
	return (*IplImage)(img)
}


/* Releases (i.e. deallocates) IPL image header */
//CVAPI(void)  cvReleaseImageHeader( IplImage** image );

/* Releases IPL image header and data */
func ReleaseImage(img *IplImage) {
	img_c := (*C.IplImage)(img)
	C.cvReleaseImage(&img_c)
}
//CVAPI(void)  cvReleaseImage( IplImage** image );

/******** matrix iterator: used for n-ary operations on dense arrays *********/

/* Returns width and height of array in elements */
func GetSizeWidth(img *IplImage) int {
	size := C.cvGetSize(unsafe.Pointer(img))
	w := int(size.width)
	return w
}
func GetSizeHeight(img *IplImage) int {
	size := C.cvGetSize(unsafe.Pointer(img))
	w := int(size.height)
	return w
}
//CVAPI(CvSize) cvGetSize( const CvArr* arr );


/* Copies source array to destination array */
func Copy(src, dst, mask *IplImage) {
	C.cvCopy(unsafe.Pointer(src), unsafe.Pointer(dst), unsafe.Pointer(mask))
}
//CVAPI(void)  cvCopy( const CvArr* src, CvArr* dst,
//                     const CvArr* mask CV_DEFAULT(NULL) );

/* Clears all the array elements (sets them to 0) */
func Zero(img *IplImage) {
	C.cvSetZero(unsafe.Pointer(img))
}
//CVAPI(void)  cvSetZero( CvArr* arr );
//#define cvZero  cvSetZero


/****************************************************************************************\
*                   Arithmetic, logic and comparison operations               *
\****************************************************************************************/


/* dst(idx) = ~src(idx) */
func Not(src, dst *IplImage) {
	C.cvNot(unsafe.Pointer(src), unsafe.Pointer(dst))
}
//CVAPI(void) cvNot( const CvArr* src, CvArr* dst );

/****************************************************************************************\
*                                Math operations                              *
\****************************************************************************************/


/****************************************************************************************\
*                                Matrix operations                            *
\****************************************************************************************/



/****************************************************************************************\
*                                    Array Statistics                         *
\****************************************************************************************/



/****************************************************************************************\
*                      Discrete Linear Transforms and Related Functions       *
\****************************************************************************************/




/****************************************************************************************\
*                              Dynamic data structures                        *
\****************************************************************************************/



/****************************************************************************************\
*                                     Drawing                                 *
\****************************************************************************************/



/****************************************************************************************\
*                                    System functions                         *
\****************************************************************************************/


/****************************************************************************************\
*                                    Data Persistence                         *
\****************************************************************************************/







